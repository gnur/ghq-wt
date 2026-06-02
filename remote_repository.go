package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gnur/ghq-wt/cmdutil"
)

// A RemoteRepository represents a remote repository.
type RemoteRepository interface {
	// URL returns the repository URL.
	URL() *url.URL
	// IsValid checks if the URL is valid.
	IsValid() bool
	// VCS returns the VCS backend that hosts the repository.
	VCS() (*VCSBackend, *url.URL, error)
}

// A GitHubRepository represents a GitHub repository. Implements RemoteRepository.
type GitHubRepository struct {
	url *url.URL
}

// URL returns URL of the repository
func (repo *GitHubRepository) URL() *url.URL {
	return repo.url
}

// IsValid determine if the repository is valid or not
func (repo *GitHubRepository) IsValid() bool {
	if strings.HasPrefix(repo.url.Path, "/blog/") {
		return false
	}
	pathComponents := strings.Split(strings.Trim(repo.url.Path, "/"), "/")
	return len(pathComponents) >= 2
}

// VCS returns VCSBackend of the repository
func (repo *GitHubRepository) VCS() (*VCSBackend, *url.URL, error) {
	u := *repo.url
	pathComponents := strings.Split(strings.Trim(strings.TrimSuffix(u.Path, ".git"), "/"), "/")
	path := "/" + strings.Join(pathComponents[0:2], "/")
	if strings.HasSuffix(u.String(), ".git") {
		path += ".git"
	}
	u.Path = path
	return GitBackend, &u, nil
}

// A GitHubGistRepository represents a GitHub Gist repository.
type GitHubGistRepository struct {
	url *url.URL
}

// URL returns URL of the GistRepository
func (repo *GitHubGistRepository) URL() *url.URL {
	return repo.url
}

// IsValid determine if the gist repository is valid or not
func (repo *GitHubGistRepository) IsValid() bool {
	return true
}

// VCS returns VCSBackend of the gist
func (repo *GitHubGistRepository) VCS() (*VCSBackend, *url.URL, error) {
	return GitBackend, repo.URL(), nil
}

// OtherRepository represents other repository
type OtherRepository struct {
	url *url.URL
}

// URL returns URL of the repository
func (repo *OtherRepository) URL() *url.URL {
	return repo.url
}

// IsValid determine if the repository is valid or not
func (repo *OtherRepository) IsValid() bool {
	return true
}

// VCS detects VCSBackend of the OtherRepository
func (repo *OtherRepository) VCS() (*VCSBackend, *url.URL, error) {
	// Detect VCS backend
	if repo.url.Scheme == "ssh" && repo.url.User.Username() == "git" {
		return GitBackend, repo.URL(), nil
	}

	if cmdutil.RunSilently("git", "ls-remote", repo.url.String()) == nil {
		return GitBackend, repo.URL(), nil
	}

	vcsStr, repoURL, err := detectGoImport(repo.url)
	if err == nil {
		if backend, ok := vcsRegistry[vcsStr]; ok {
			return backend, repoURL, nil
		}
	}

	return nil, nil, fmt.Errorf("unsupported VCS, url=%s: could not detect git repository", repo.URL())
}

// NewRemoteRepository returns new RemoteRepository object from URL
func NewRemoteRepository(u *url.URL) (RemoteRepository, error) {
	repo := func() RemoteRepository {
		switch u.Host {
		case "github.com":
			return &GitHubRepository{u}
		case "gist.github.com":
			return &GitHubGistRepository{u}
		default:
			return &OtherRepository{u}
		}
	}()
	if !repo.IsValid() {
		return nil, fmt.Errorf("not a valid repository: %s", u)
	}
	return repo, nil
}
