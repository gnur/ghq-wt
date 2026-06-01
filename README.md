# ghq

[![Build Status](https://github.com/gnur/ghq-wt/workflows/test/badge.svg?branch=master)](https://github.com/gnur/ghq-wt/actions?workflow=test)
[![Coverage](https://codecov.io/gh/gnur/ghq-wt/branch/master/graph/badge.svg)](https://codecov.io/gh/gnur/ghq-wt)

## Name

ghq - Manage remote repository clones

## Description

`ghq` provides a way to organize remote repository clones, like `go get` does. When you clone a remote repository by `ghq get`, ghq makes a directory under a specific root directory (by default `~/ghq`) using the remote repository URL's host and path.

Git repositories are cloned using a **worktree layout** by default: a bare clone is placed in `.bare/` and the default branch is checked out as a git worktree. This allows you to have multiple branches checked out simultaneously under the same repository path.

```
$ ghq get https://github.com/gnur/ghq-wt
# Creates:
#   ~/ghq/github.com/gnur/ghq-wt/.bare/   (bare clone)
#   ~/ghq/github.com/gnur/ghq-wt/main/    (worktree for default branch)
```

You can also list local repositories (`ghq list`).

## Synopsis

```
ghq get [-u] [-p] [--shallow] [--vcs <vcs>] [--look] [--silent] [--branch] [--no-recursive] [--bare] [--partial blobless|treeless] <repository URL>|<host>/<user>/<project>|<user>/<project>|<project>
ghq list [-p] [-e] [<query>]
ghq create [--vcs <vcs>] <repository URL>|<host>/<user>/<project>|<user>/<project>|<project>
ghq rm [--dry-run] <repository URL>|<host>/<user>/<project>|<user>/<project>|<project>
ghq migrate [-y] [--dry-run] <local repository path>
ghq migrate-worktree [--dry-run] <repository>
ghq root [--all]
```

## Commands

### get

Clone a remote repository under ghq root directory (see [Directory Structures](#directory-structures) below). `ghq clone` is an alias for this command.

If the repository is already cloned to local, nothing will happen unless `-u` (`--update`) flag is supplied, in which case the local repository is updated (fetch in the bare repo).

When you use `-p` option, the repository is cloned via SSH protocol.

If there are multiple `ghq.root`s, existing local clones are searched first. Then a new repository clone is created under the primary root if none is found.

**Options:**

- `--shallow` — Perform a shallow clone (Git only, `--depth 1`). Be careful that a shallow-cloned repository cannot be pushed to remote.
- `--branch`, `-b` — Clone with specified branch. The worktree will be named after this branch. Supported for Git, Mercurial, Subversion and git-svn.
- `--no-recursive` — Prevent recursive fetching of submodules.
- `--bare` — Perform a plain bare clone (skips the worktree layout, creates a traditional bare repo).
- `--partial` — Perform a partial clone (`blobless` for `--filter=blob:none`, `treeless` for `--filter=tree:0`).

### list

List locally cloned repositories. If a query argument is given, only repositories whose names contain that query text are listed.

- `-e`, `--exact` — Forces the match to be exact (query equals _project_, _user/project_ or _host/user/project_).
- `-p`, `--full-path` — Print full paths to the repository root instead of relative ones.
- `--unique` — Print unique subpaths.

### root

Prints repositories' root (i.e. `ghq.root`). Without `--all` option, the primary one is shown.

### rm

Remove local repository. If `--dry-run` option is given, the repository is not actually removed but the path to it is printed. Handles worktree-layout repos correctly (removes `.bare/` and all worktrees).

### create

Creates a new repository.

### migrate

Migrate an existing repository directory to the ghq-managed directory structure. The command detects the VCS backend, retrieves the remote URL, and moves the repository to the appropriate location under ghq root.

### migrate-worktree

Convert an existing git clone from standard layout (`.git` directory) to the worktree layout (`.bare/` + branch worktrees). The current branch becomes the first worktree.

```
$ ghq migrate-worktree github.com/gnur/ghq-wt
# Converts:
#   .git/           -> .bare/
#   working files   -> main/   (or whatever the current branch is)
```

Use `--dry-run` to preview the migration without making changes.

## Worktree Layout

For Git repositories, `ghq get` uses a worktree-based layout by default:

```
~/ghq/github.com/gnur/ghq-wt/
├── .bare/          # bare clone (shared object store)
└── main/           # git worktree (default branch)
```

This allows multiple branches to be checked out simultaneously:

```
~/ghq/github.com/gnur/ghq-wt/
├── .bare/
├── main/
└── feature-xyz/    # additional worktree
```

**Key behaviors:**

- `ghq get` clones bare and creates a worktree for the default (or specified) branch
- `ghq list` shows the worktree directories (e.g. `github.com/gnur/ghq-wt/main`)
- `ghq get --update` fetches in the bare repo
- `ghq get --bare` bypasses the worktree layout and does a traditional bare clone
- Non-git VCS backends (hg, svn, bzr, etc.) clone as before (no worktree support)

To convert an existing standard clone to the worktree layout:

```
ghq migrate-worktree github.com/gnur/ghq-wt
```

## Configuration

Configuration uses `git-config` variables.

### ghq.root

The path to directory under which cloned repositories are placed. See [Directory Structures](#directory-structures) below. Defaults to `~/ghq`.

This variable can have multiple values. If so, the last one becomes the primary one (i.e. new repository clones are always created under it). You may want to specify `$GOPATH/src` as a secondary root (environment variables should be expanded).

### ghq.user

When specifying only the repository name without slashes as in `ghq get <project>`, ghq attempts to auto-complete the repository owner. By default, the owner used is the value of the environment variable `USER` (or `USERNAME` on Windows). Setting this option allows you to explicitly specify the owner.

### ghq.completeUser

Rather than always using your own username for owner completion, you may want to complete the owner with the same name as the repository. For example, fetch `ruby` as `github.com/ruby/ruby`, `vim` as `github.com/vim/vim`, and `peco` as `github.com/peco/peco`. If you prefer this behavior, set this option to `false`.

### ghq.defaultHost

The default host used when the repository specification omits the host. For example, `ghq get owner/project` normally resolves to `github.com/owner/project`. If this option is set, the specified host will be used instead.

### ghq.\<url\>.vcs

Explicitly specify the VCS for a remote repository. The URL is matched against `<url>` using `git config --get-urlmatch`.

Accepted values: `git`, `github` (alias for git), `subversion`, `svn`, `git-svn`, `mercurial`, `hg`, `darcs`, `fossil`, `bazaar`, `bzr`.

Requires Git 1.8.5 or higher.

### ghq.\<url\>.root

Specify a repository-specific root directory instead of the common ghq root directory. The URL is matched against `<url>` using `git config --get-urlmatch`.

### Example configuration (.gitconfig)

```ini
[ghq "https://git.example.com/repos/"]
vcs = git
root = ~/myproj
```

## Environment Variables

### GHQ_ROOT

If set to a path, this value is used as the only root directory regardless of other existing `ghq.root` settings.

## Directory Structures

Local repositories are placed under `ghq.root` using the URL structure:

```
~/ghq
├── code.google.com/
│   └── p/
│       └── vim/
└── github.com/
    ├── google/
    │   └── go-github/
    │       └── main/
    ├── motemen/
    │   └── ghq/
    │       ├── .bare/
    │       └── main/
    └── urfave/
        └── cli/
            └── main/
```

## Installation

### macOS

```
brew install ghq
```

### Void Linux

```
xbps-install -S ghq
```

### GNU Guix

```
guix install ghq
```

### Windows + scoop

```
scoop install ghq
```

### go install

```
go install github.com/gnur/ghq-wt@latest
```

### conda

```
conda install -c conda-forge go-ghq
```

### [asdf-vm](https://github.com/asdf-vm/asdf)

```
asdf plugin add ghq
asdf install ghq latest
```

### [mise-en-place](https://github.com/jdx/mise)

```
mise install ghq
mise use ghq
```

### Build from source

```
git clone https://github.com/gnur/ghq-wt .
make install
```

Built binaries are available from [GitHub Releases](https://github.com/gnur/ghq-wt/releases).

## Handbook

You can buy the "ghq-handbook" from [Leanpub](https://leanpub.com/ghq-handbook) for more detailed usage.

The source Markdown files are also available for free from [github.com/Songmu/ghq-handbook](https://github.com/Songmu/ghq-handbook).

Currently only a Japanese version is available. Translations are welcome!

## Authors

- motemen <motemen@gmail.com> — [Sponsor](https://github.com/sponsors/motemen)
- Songmu <y.songmu@gmail.com> — [Sponsor](https://github.com/sponsors/Songmu)
