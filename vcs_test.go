package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/gnur/ghq-wt/cmdutil"
)

var remoteDummyURL = mustParseURL("https://example.com/git/repo")

func TestVCSBackend(t *testing.T) {
	tempDir := newTempDir(t)
	localDir := filepath.Join(tempDir, "repo")
	var _commands []*exec.Cmd
	lastCommand := func() *exec.Cmd { return _commands[len(_commands)-1] }
	defer func(orig func(cmd *exec.Cmd) error) {
		cmdutil.CommandRunner = orig
	}(cmdutil.CommandRunner)
	cmdutil.CommandRunner = func(cmd *exec.Cmd) error {
		_commands = append(_commands, cmd)
		return nil
	}

	testCases := []struct {
		name   string
		f      func() error
		expect []string
		dir    string
	}{{
		name: "[git] clone (worktree layout)",
		f: func() error {
			return GitBackend.Clone(&vcsGetOption{
				url: remoteDummyURL,
				dir: localDir,
			})
		},
		// Last command is worktree add (bare clone happens first)
		expect: []string{"git", "worktree", "add", filepath.Join(localDir, "main"), "main"},
		dir:    filepath.Join(localDir, ".bare"),
	}, {
		name: "[git] shallow clone (worktree layout)",
		f: func() error {
			return GitBackend.Clone(&vcsGetOption{
				url:     remoteDummyURL,
				dir:     localDir,
				shallow: true,
				silent:  true,
			})
		},
		expect: []string{"git", "worktree", "add", filepath.Join(localDir, "main"), "main"},
		dir:    filepath.Join(localDir, ".bare"),
	}, {
		name: "[git] clone specific branch (worktree layout)",
		f: func() error {
			return GitBackend.Clone(&vcsGetOption{
				url:    remoteDummyURL,
				dir:    localDir,
				branch: "hello",
			})
		},
		expect: []string{"git", "worktree", "add", filepath.Join(localDir, "hello"), "hello"},
		dir:    filepath.Join(localDir, ".bare"),
	}, {
		name: "[git] update",
		f: func() error {
			return GitBackend.Update(&vcsGetOption{
				dir: localDir,
			})
		},
		expect: []string{"git", "pull", "--ff-only"},
		dir:    localDir,
	}, {
		name: "[git] fetch",
		f: func() error {
			defer func(orig func(cmd *exec.Cmd) error) {
				cmdutil.CommandRunner = orig
			}(cmdutil.CommandRunner)
			cmdutil.CommandRunner = func(cmd *exec.Cmd) error {
				_commands = append(_commands, cmd)
				if reflect.DeepEqual(cmd.Args, []string{"git", "rev-parse", "@{upstream}"}) {
					return fmt.Errorf("[test] failed to git rev-parse @{upstream}")
				}
				return nil
			}
			return GitBackend.Update(&vcsGetOption{
				dir: localDir,
			})
		},
		expect: []string{"git", "fetch"},
		dir:    localDir,
	}, {
		name: "[git] recursive (worktree layout)",
		f: func() error {
			return GitBackend.Clone(&vcsGetOption{
				url:       remoteDummyURL,
				dir:       localDir,
				recursive: true,
			})
		},
		// Last command is submodule update in the worktree dir
		expect: []string{"git", "submodule", "update", "--init", "--recursive"},
		dir:    filepath.Join(localDir, "main"),
	}, {
		name: "[git] update recursive",
		f: func() error {
			return GitBackend.Update(&vcsGetOption{
				dir:       localDir,
				recursive: true,
			})
		},
		expect: []string{"git", "submodule", "update", "--init", "--recursive"},
		dir:    localDir,
	}, {
		name: "[git] bare clone (explicit --bare, no worktree layout)",
		f: func() error {
			return GitBackend.Clone(&vcsGetOption{
				url:    remoteDummyURL,
				dir:    localDir,
				bare:   true,
				silent: true,
			})
		},
		expect: []string{"git", "clone", "--bare", remoteDummyURL.String(), localDir},
	}, {
		name: "[git] (partial) blobless clone (worktree layout)",
		f: func() error {
			return GitBackend.Clone(&vcsGetOption{
				url:     remoteDummyURL,
				dir:     localDir,
				partial: "blobless",
			})
		},
		expect: []string{"git", "worktree", "add", filepath.Join(localDir, "main"), "main"},
		dir:    filepath.Join(localDir, ".bare"),
	}, {
		name: "[git] (partial) treeless clone (worktree layout)",
		f: func() error {
			return GitBackend.Clone(&vcsGetOption{
				url:     remoteDummyURL,
				dir:     localDir,
				partial: "treeless",
			})
		},
		expect: []string{"git", "worktree", "add", filepath.Join(localDir, "main"), "main"},
		dir:    filepath.Join(localDir, ".bare"),
	}}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if err := tc.f(); err != nil {
				t.Errorf("error should be nil, but: %s", err)
			}
			c := lastCommand()
			if !reflect.DeepEqual(c.Args, tc.expect) {
				t.Errorf("\ngot:    %+v\nexpect: %+v", c.Args, tc.expect)
			}
			if c.Dir != tc.dir {
				t.Errorf("got: %s, expect: %s", c.Dir, tc.dir)
			}
		})
	}
}
