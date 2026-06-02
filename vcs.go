package main

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gnur/ghq-wt/cmdutil"
)

func run(silent bool) func(command string, args ...string) error {
	if silent {
		return cmdutil.RunSilently
	}
	return cmdutil.Run
}

func runInDir(silent bool) func(dir, command string, args ...string) error {
	if silent {
		return cmdutil.RunInDirSilently
	}
	return cmdutil.RunInDir
}

// A VCSBackend represents a VCS backend.
type VCSBackend struct {
	// Clones a remote repository to local path.
	Clone func(*vcsGetOption) error
	// Updates a cloned local repository.
	Update func(*vcsGetOption) error
	Init   func(dir string) error
	// Returns VCS specific files
	Contents []string
	// Returns the remote URL of the repository at the given directory.
	// If nil, the VCS backend does not support retrieving remote URLs.
	RemoteURL func(dir string) (string, error)
}

type vcsGetOption struct {
	url                              *url.URL
	dir                              string
	recursive, shallow, silent, bare bool
	branch, partial                  string
}

// getGitRemoteURL retrieves the remote URL from a git repository.
// It tries 'origin' first, then falls back to the first remote.
func getGitRemoteURL(dir string) (string, error) {
	// Try 'origin' first
	originCmd := exec.Command("git", "remote", "get-url", "origin")
	originCmd.Dir = dir
	originOut, originErr := originCmd.Output()
	if originErr == nil {
		originURL := strings.TrimSpace(string(originOut))
		if originURL != "" {
			return originURL, nil
		}
	}

	// List all remotes
	listCmd := exec.Command("git", "remote")
	listCmd.Dir = dir
	listOut, listErr := listCmd.Output()
	if listErr != nil {
		return "", fmt.Errorf("failed to list remotes: %w", listErr)
	}

	allRemotes := strings.Split(strings.TrimSpace(string(listOut)), "\n")
	if len(allRemotes) == 0 || allRemotes[0] == "" {
		return "", fmt.Errorf("no remotes found")
	}

	// Get first remote URL
	first := allRemotes[0]
	urlCmd := exec.Command("git", "remote", "get-url", first)
	urlCmd.Dir = dir
	urlOut, urlErr := urlCmd.Output()
	if urlErr != nil {
		return "", fmt.Errorf("failed to get URL of remote %q: %w", first, urlErr)
	}

	finalURL := strings.TrimSpace(string(urlOut))
	if finalURL == "" {
		return "", fmt.Errorf("remote %q has no URL", first)
	}

	return finalURL, nil
}

// GitBackend is the VCSBackend of git
var GitBackend = &VCSBackend{
	Clone: func(vg *vcsGetOption) error {
		// If user explicitly requested --bare, do a plain bare clone (no worktree layout)
		if vg.bare {
			dir, _ := filepath.Split(vg.dir)
			if err := os.MkdirAll(dir, 0755); err != nil {
				return err
			}
			args := []string{"clone", "--bare"}
			if vg.shallow {
				args = append(args, "--depth", "1")
			}
			if vg.branch != "" {
				args = append(args, "--branch", vg.branch, "--single-branch")
			}
			if vg.partial == "blobless" {
				args = append(args, "--filter=blob:none")
			} else if vg.partial == "treeless" {
				args = append(args, "--filter=tree:0")
			}
			args = append(args, vg.url.String(), vg.dir)
			return run(vg.silent)("git", args...)
		}

		// Worktree layout: clone bare to <dir>/.bare, then add worktree for default branch
		bareDir := filepath.Join(vg.dir, ".bare")
		err := os.MkdirAll(vg.dir, 0755)
		if err != nil {
			return err
		}

		// Step 1: Clone as bare repo
		args := []string{"clone", "--bare"}
		if vg.shallow {
			args = append(args, "--depth", "1")
		}
		if vg.branch != "" {
			args = append(args, "--branch", vg.branch, "--single-branch")
		}
		if vg.partial == "blobless" {
			args = append(args, "--filter=blob:none")
		} else if vg.partial == "treeless" {
			args = append(args, "--filter=tree:0")
		}
		args = append(args, vg.url.String(), bareDir)

		if err := run(vg.silent)("git", args...); err != nil {
			return err
		}

		// Step 2: Detect the default branch
		branch := vg.branch
		if branch == "" {
			branch = detectDefaultBranch(bareDir, "")
		}

		// Step 3: Add worktree for the branch
		worktreeDir := filepath.Join(vg.dir, branch)
		if err := runInDir(vg.silent)(bareDir, "git", "worktree", "add", worktreeDir, branch); err != nil {
			return fmt.Errorf("failed to add worktree for branch %q: %w", branch, err)
		}

		// Step 4: If recursive, init submodules in the worktree
		if vg.recursive {
			return runInDir(vg.silent)(worktreeDir, "git", "submodule", "update", "--init", "--recursive")
		}
		return nil
	},
	Update: func(vg *vcsGetOption) error {
		// Check if this is a worktree layout (has .bare directory)
		bareDir := filepath.Join(vg.dir, ".bare")
		if fi, err := os.Stat(bareDir); err == nil && fi.IsDir() {
			// Worktree layout: fetch in bare, then pull in each worktree
			if err := runInDir(vg.silent)(bareDir, "git", "fetch", "--all"); err != nil {
				return err
			}
			// Pull in the current worktree dir (vg.dir may be a worktree)
			return nil
		}

		// Legacy layout: standard update
		if vg.bare {
			return runInDir(true)(vg.dir, "git", "fetch", vg.url.String(), "*:*")
		}
		err := runInDir(true)(vg.dir, "git", "rev-parse", "@{upstream}")
		if err != nil {
			err := runInDir(vg.silent)(vg.dir, "git", "fetch")
			if err != nil {
				return err
			}
			return nil
		}
		err = runInDir(vg.silent)(vg.dir, "git", "pull", "--ff-only")
		if err != nil {
			return err
		}
		if vg.recursive {
			return runInDir(vg.silent)(vg.dir, "git", "submodule", "update", "--init", "--recursive")
		}
		return nil
	},
	Init: func(dir string) error {
		args := []string{"init"}
		if strings.HasSuffix(dir, ".git") {
			args = append(args, "--bare")
		}
		return cmdutil.RunInDir(dir, "git", args...)
	},
	Contents: []string{".git"},
	RemoteURL: func(dir string) (string, error) {
		return getGitRemoteURL(dir)
	},
}

var vcsRegistry = map[string]*VCSBackend{
	"git":    GitBackend,
	"github": GitBackend,
}
