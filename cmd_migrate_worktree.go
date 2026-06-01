package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v3"
	"github.com/gnur/ghq-wt/logger"
)

func doMigrateWorktree(ctx context.Context, cmd *cli.Command) error {
	var (
		name   = cmd.Args().First()
		dry    = cmd.Bool("dry-run")
		w      = cmd.Root().Writer
		bare   = cmd.Bool("bare")
	)

	if name == "" {
		return fmt.Errorf("repository name or path is required")
	}

	// Resolve the repository path
	var repoPath string
	if filepath.IsAbs(name) {
		repoPath = name
	} else {
		u, err := newURL(name, false, true)
		if err != nil {
			return err
		}
		localRepo, err := LocalRepositoryFromURL(u, bare)
		if err != nil {
			return err
		}
		repoPath = localRepo.FullPath
	}

	// Verify it's an existing git repo (standard layout with .git directory)
	gitDir := filepath.Join(repoPath, ".git")
	fi, err := os.Stat(gitDir)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("%s is not a git repository (no .git found)", repoPath)
		}
		return err
	}
	if !fi.IsDir() {
		return fmt.Errorf("%s already appears to be a worktree (has .git file, not directory)", repoPath)
	}

	// Check if already in worktree layout
	bareDir := filepath.Join(repoPath, ".bare")
	if _, err := os.Stat(bareDir); err == nil {
		return fmt.Errorf("%s is already in worktree layout (.bare/ exists)", repoPath)
	}

	// Detect current branch
	branchCmd := exec.Command("git", "symbolic-ref", "--short", "HEAD")
	branchCmd.Dir = repoPath
	branchOut, err := branchCmd.Output()
	if err != nil {
		return fmt.Errorf("failed to detect current branch: %w (are you in detached HEAD?)", err)
	}
	currentBranch := strings.TrimSpace(string(branchOut))

	// Check for uncommitted changes
	statusCmd := exec.Command("git", "status", "--porcelain")
	statusCmd.Dir = repoPath
	statusOut, err := statusCmd.Output()
	if err != nil {
		return fmt.Errorf("failed to check git status: %w", err)
	}
	if len(statusOut) > 0 {
		return fmt.Errorf("repository has uncommitted changes; commit or stash them first")
	}

	worktreeDir := filepath.Join(repoPath, currentBranch)

	if dry {
		fmt.Fprintf(w, "Would migrate %s to worktree layout:\n", repoPath)
		fmt.Fprintf(w, "  .git/ -> .bare/\n")
		fmt.Fprintf(w, "  working files -> %s/\n", currentBranch)
		fmt.Fprintf(w, "  branch: %s\n", currentBranch)
		return nil
	}

	logger.Log("migrate", fmt.Sprintf("%s -> worktree layout (branch: %s)", repoPath, currentBranch))

	// Step 1: Rename .git to .bare
	if err := os.Rename(gitDir, bareDir); err != nil {
		return fmt.Errorf("failed to rename .git to .bare: %w", err)
	}

	// Step 2: Mark it as a bare repo
	cfgCmd := exec.Command("git", "config", "core.bare", "true")
	cfgCmd.Dir = bareDir
	if err := cfgCmd.Run(); err != nil {
		// Try to undo
		os.Rename(bareDir, gitDir)
		return fmt.Errorf("failed to set core.bare=true: %w", err)
	}

	// Step 3: Create worktree directory and move all working files there
	if err := os.MkdirAll(worktreeDir, 0755); err != nil {
		os.Rename(bareDir, gitDir)
		return fmt.Errorf("failed to create worktree directory: %w", err)
	}

	// Move all files/dirs (except .bare and the new worktree dir) into worktree dir
	entries, err := os.ReadDir(repoPath)
	if err != nil {
		os.Rename(bareDir, gitDir)
		return fmt.Errorf("failed to list repo directory: %w", err)
	}

	for _, entry := range entries {
		name := entry.Name()
		if name == ".bare" || name == currentBranch {
			continue
		}
		src := filepath.Join(repoPath, name)
		dst := filepath.Join(worktreeDir, name)
		if err := os.Rename(src, dst); err != nil {
			return fmt.Errorf("failed to move %s to worktree: %w", name, err)
		}
	}

	// Step 4: Create .git file in worktree pointing to .bare/worktrees/<branch>
	// First, register the worktree with git
	// We need to unset core.bare temporarily to add the worktree entry
	cfgCmd = exec.Command("git", "config", "core.bare", "false")
	cfgCmd.Dir = bareDir
	cfgCmd.Run() // best effort

	// Write .git file pointing to the bare repo's worktree tracking dir
	worktreesTrackingDir := filepath.Join(bareDir, "worktrees", currentBranch)
	if err := os.MkdirAll(worktreesTrackingDir, 0755); err != nil {
		return fmt.Errorf("failed to create worktree tracking dir: %w", err)
	}

	// Write the gitdir file in the worktree tracking dir
	gitdirContent := filepath.Join(worktreeDir, ".git") + "\n"
	if err := os.WriteFile(filepath.Join(worktreesTrackingDir, "gitdir"), []byte(gitdirContent), 0644); err != nil {
		return fmt.Errorf("failed to write gitdir file: %w", err)
	}

	// Write commondir in worktree tracking dir
	commondir := filepath.Join("..", "..")
	if err := os.WriteFile(filepath.Join(worktreesTrackingDir, "commondir"), []byte(commondir+"\n"), 0644); err != nil {
		return fmt.Errorf("failed to write commondir: %w", err)
	}

	// Write HEAD in worktree tracking dir
	headContent := fmt.Sprintf("ref: refs/heads/%s\n", currentBranch)
	if err := os.WriteFile(filepath.Join(worktreesTrackingDir, "HEAD"), []byte(headContent), 0644); err != nil {
		return fmt.Errorf("failed to write worktree HEAD: %w", err)
	}

	// Write .git file in the worktree working dir
	dotGitContent := fmt.Sprintf("gitdir: %s\n", worktreesTrackingDir)
	if err := os.WriteFile(filepath.Join(worktreeDir, ".git"), []byte(dotGitContent), 0644); err != nil {
		return fmt.Errorf("failed to write .git file in worktree: %w", err)
	}

	// Set core.bare back to true
	cfgCmd = exec.Command("git", "config", "core.bare", "true")
	cfgCmd.Dir = bareDir
	cfgCmd.Run()

	fmt.Fprintf(w, "Migrated %s to worktree layout\n", repoPath)
	fmt.Fprintf(w, "  bare repo: %s\n", bareDir)
	fmt.Fprintf(w, "  worktree:  %s\n", worktreeDir)
	return nil
}
