package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestGitWorktreeClone(t *testing.T) {
	// Create a fake remote bare repo to clone from
	remoteDir := newTempDir(t)
	remoteRepo := filepath.Join(remoteDir, "origin.git")

	// Init bare remote
	if err := exec.Command("git", "init", "--bare", remoteRepo).Run(); err != nil {
		t.Fatalf("failed to init bare remote: %v", err)
	}

	// Create a temporary working clone to push a commit
	workDir := filepath.Join(remoteDir, "work")
	if err := exec.Command("git", "clone", remoteRepo, workDir).Run(); err != nil {
		t.Fatalf("failed to clone for setup: %v", err)
	}

	// Create a commit on "main" branch
	cmd := exec.Command("git", "checkout", "-b", "main")
	cmd.Dir = workDir
	cmd.Run()

	cmd = exec.Command("git", "commit", "--allow-empty", "-m", "initial")
	cmd.Dir = workDir
	cmd.Env = append(os.Environ(),
		"GIT_AUTHOR_NAME=Test",
		"GIT_AUTHOR_EMAIL=test@test.com",
		"GIT_COMMITTER_NAME=Test",
		"GIT_COMMITTER_EMAIL=test@test.com",
	)
	if err := cmd.Run(); err != nil {
		t.Fatalf("failed to create commit: %v", err)
	}

	cmd = exec.Command("git", "push", "origin", "main")
	cmd.Dir = workDir
	if err := cmd.Run(); err != nil {
		t.Fatalf("failed to push: %v", err)
	}

	// Set HEAD to main in the bare remote
	cmd = exec.Command("git", "symbolic-ref", "HEAD", "refs/heads/main")
	cmd.Dir = remoteRepo
	cmd.Run()

	// Now test our worktree clone
	destDir := filepath.Join(remoteDir, "dest")
	if err := os.MkdirAll(destDir, 0755); err != nil {
		t.Fatal(err)
	}

	remoteURL := mustParseURL("file://" + remoteRepo)
	err := GitBackend.Clone(&vcsGetOption{
		url: remoteURL,
		dir: destDir,
	})
	if err != nil {
		t.Fatalf("GitBackend.Clone failed: %v", err)
	}

	// Verify .bare directory exists
	bareDir := filepath.Join(destDir, ".bare")
	if fi, err := os.Stat(bareDir); err != nil || !fi.IsDir() {
		t.Errorf(".bare directory should exist at %s", bareDir)
	}

	// Verify worktree directory exists
	worktreeDir := filepath.Join(destDir, "main")
	if fi, err := os.Stat(worktreeDir); err != nil || !fi.IsDir() {
		t.Errorf("worktree directory should exist at %s", worktreeDir)
	}

	// Verify .git file in worktree (not directory)
	dotGit := filepath.Join(worktreeDir, ".git")
	fi, err := os.Lstat(dotGit)
	if err != nil {
		t.Fatalf(".git should exist in worktree: %v", err)
	}
	if fi.IsDir() {
		t.Error(".git in worktree should be a file, not a directory")
	}

	// Verify we can run git commands in the worktree
	cmd = exec.Command("git", "status")
	cmd.Dir = worktreeDir
	if err := cmd.Run(); err != nil {
		t.Errorf("git status failed in worktree: %v", err)
	}

	// Verify branch is correct
	cmd = exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Dir = worktreeDir
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("failed to get branch: %v", err)
	}
	if got := string(out[:len(out)-1]); got != "main" {
		t.Errorf("expected branch 'main', got %q", got)
	}
}

func TestGitWorktreeCloneWithBranch(t *testing.T) {
	// Create a fake remote with a specific branch
	remoteDir := newTempDir(t)
	remoteRepo := filepath.Join(remoteDir, "origin.git")

	if err := exec.Command("git", "init", "--bare", remoteRepo).Run(); err != nil {
		t.Fatalf("failed to init bare remote: %v", err)
	}

	workDir := filepath.Join(remoteDir, "work")
	if err := exec.Command("git", "clone", remoteRepo, workDir).Run(); err != nil {
		t.Fatalf("failed to clone for setup: %v", err)
	}

	cmd := exec.Command("git", "checkout", "-b", "develop")
	cmd.Dir = workDir
	cmd.Run()

	cmd = exec.Command("git", "commit", "--allow-empty", "-m", "initial")
	cmd.Dir = workDir
	cmd.Env = append(os.Environ(),
		"GIT_AUTHOR_NAME=Test",
		"GIT_AUTHOR_EMAIL=test@test.com",
		"GIT_COMMITTER_NAME=Test",
		"GIT_COMMITTER_EMAIL=test@test.com",
	)
	cmd.Run()

	cmd = exec.Command("git", "push", "origin", "develop")
	cmd.Dir = workDir
	cmd.Run()

	// Clone with specific branch
	destDir := filepath.Join(remoteDir, "dest")
	os.MkdirAll(destDir, 0755)

	remoteURL := mustParseURL("file://" + remoteRepo)
	err := GitBackend.Clone(&vcsGetOption{
		url:    remoteURL,
		dir:    destDir,
		branch: "develop",
	})
	if err != nil {
		t.Fatalf("GitBackend.Clone failed: %v", err)
	}

	// Verify worktree is named after the branch
	worktreeDir := filepath.Join(destDir, "develop")
	if fi, err := os.Stat(worktreeDir); err != nil || !fi.IsDir() {
		t.Errorf("worktree directory 'develop' should exist at %s", worktreeDir)
	}
}

func TestGitBareCloneNoWorktree(t *testing.T) {
	// When --bare is explicitly passed, should NOT use worktree layout
	remoteDir := newTempDir(t)
	remoteRepo := filepath.Join(remoteDir, "origin.git")

	if err := exec.Command("git", "init", "--bare", remoteRepo).Run(); err != nil {
		t.Fatalf("failed to init bare remote: %v", err)
	}

	destDir := filepath.Join(remoteDir, "dest.git")

	remoteURL := mustParseURL("file://" + remoteRepo)
	err := GitBackend.Clone(&vcsGetOption{
		url:  remoteURL,
		dir:  destDir,
		bare: true,
	})
	if err != nil {
		t.Fatalf("GitBackend.Clone failed: %v", err)
	}

	// Should NOT have .bare subdirectory (it's a plain bare clone)
	if _, err := os.Stat(filepath.Join(destDir, ".bare")); !os.IsNotExist(err) {
		t.Error(".bare directory should NOT exist for explicit --bare clone")
	}

	// Should have HEAD (bare repo indicator)
	if _, err := os.Stat(filepath.Join(destDir, "HEAD")); err != nil {
		t.Error("HEAD should exist in bare clone")
	}
}

func TestDetectDefaultBranch(t *testing.T) {
	// Create a bare repo with a known HEAD
	tempDir := newTempDir(t)
	bareRepo := filepath.Join(tempDir, "test.git")

	if err := exec.Command("git", "init", "--bare", bareRepo).Run(); err != nil {
		t.Fatalf("failed to init bare repo: %v", err)
	}

	// Set HEAD to refs/heads/develop
	cmd := exec.Command("git", "symbolic-ref", "HEAD", "refs/heads/develop")
	cmd.Dir = bareRepo
	if err := cmd.Run(); err != nil {
		t.Fatalf("failed to set HEAD: %v", err)
	}

	branch := detectDefaultBranch(bareRepo, "")
	if branch != "develop" {
		t.Errorf("expected 'develop', got %q", branch)
	}
}

func TestDetectDefaultBranchFallback(t *testing.T) {
	// Non-existent dir should fall back to "main"
	branch := detectDefaultBranch("/nonexistent", "")
	if branch != "main" {
		t.Errorf("expected 'main' fallback, got %q", branch)
	}
}
