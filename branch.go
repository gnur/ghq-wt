package main

import (
	"os/exec"
	"strings"
)

// detectDefaultBranch queries a remote or local bare repo for its default branch.
// It first tries reading HEAD from a local bare repo path (if it exists),
// then falls back to `git ls-remote --symref <url> HEAD`.
// Returns "main" as ultimate fallback.
func detectDefaultBranch(bareDir string, remoteURL string) string {
	// Try reading HEAD from local bare repo first (avoids network call)
	if bareDir != "" {
		cmd := exec.Command("git", "symbolic-ref", "HEAD")
		cmd.Dir = bareDir
		out, err := cmd.Output()
		if err == nil {
			ref := strings.TrimSpace(string(out))
			if branch, ok := strings.CutPrefix(ref, "refs/heads/"); ok {
				return branch
			}
		}
	}

	// Fall back to ls-remote
	if remoteURL != "" {
		cmd := exec.Command("git", "ls-remote", "--symref", remoteURL, "HEAD")
		out, err := cmd.Output()
		if err == nil {
			// Output format: "ref: refs/heads/main\tHEAD\n..."
			for _, line := range strings.Split(string(out), "\n") {
				if strings.HasPrefix(line, "ref: refs/heads/") {
					parts := strings.Fields(line)
					if len(parts) >= 2 {
						branch := strings.TrimPrefix(parts[0], "ref: refs/heads/")
						return branch
					}
				}
			}
		}
	}

	return "main"
}
