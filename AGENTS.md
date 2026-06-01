# AGENTS.md

## Overview

`ghq` is a CLI tool for managing local clones of remote repositories, organized under a root directory (like `go get` but for any VCS). Written in Go, uses `urfave/cli/v3`.

Module path: `github.com/gnur/ghq-wt`

## Commands

```bash
go build ./...          # build
go test ./...           # test (all packages: root, cmdutil, logger)
go test -run TestFoo    # single test
staticcheck ./...       # lint (install: go install honnef.co/go/tools/cmd/staticcheck@latest)
```

Build with version info: `go build -ldflags="-s -w -X main.revision=$(git rev-parse --short HEAD)"`

## Package structure

- Root package (`main`): CLI commands (`cmd_*.go`), repository types, URL parsing, VCS backends, getter logic
- `cmdutil/`: subprocess execution helpers (`Run`, `RunInDir`, `RunSilently`)
- `logger/`: logging wrapper

There are no submodules or workspace files — it's a single-module repo with 3 packages.

## Architecture notes

- Commands are defined in `commands.go` (registration) and implemented in `cmd_*.go` files
- Repository types (`GitHubRepository`, `ChiselRepository`, etc.) live in `remote_repository.go` — they share an interface but are not in separate packages
- VCS backends (git, hg, svn, bzr, etc.) are in `vcs.go` with a `VCSBackend` interface
- `getter.go` orchestrates clone/update logic
- Platform-specific code uses build tags: `helpers_unix.go`, `helpers_windows.go`
- `go_import.go` handles Go import path meta-tag resolution

## Testing

- Tests run on Linux, macOS, and Windows in CI
- Many tests create temp directories and fake git backends (`withFakeGitBackend` helper in test files)
- `helpers_test.go` has shared test utilities: `newTempDir`, `setEnv`, `mustParseURL`
- No external services required; tests are self-contained
- `-short` flag works and tests complete in ~5s

## CI

- PR/push: `go test -coverprofile` on matrix (ubuntu/macos/windows)
- Lint: `reviewdog` with staticcheck
- Release: `godzil` + `ghr` (tag-triggered)
