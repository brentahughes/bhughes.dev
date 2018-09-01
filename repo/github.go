package repo

import (
	"context"

	"github.com/google/go-github/github"
)

type GithubClient struct {
	client   *github.Client
	username string
}

type GithubRepo struct {
	Name         string
	URL          string
	Contribution bool
}

func (r *GithubRepo) GetName() string {
	return r.Name
}

func (r *GithubRepo) GetURL() string {
	return r.URL
}

func (r *GithubRepo) GetRepo() string {
	return "github"
}

func (r *GithubRepo) IsContribution() bool {
	return r.Contribution
}

func (c *GithubClient) GetPublicRepos() ([]Repo, error) {
	r, _, err := c.client.Repositories.List(context.Background(), c.username, nil)
	if err != nil {
		return nil, err
	}

	repos := make([]Repo, 0)
	for _, repo := range r {
		repos = append(repos, &GithubRepo{
			Name:         *repo.Name,
			URL:          *repo.SVNURL,
			Contribution: *repo.Fork,
		})
	}

	return repos, nil
}

func (c *GithubClient) GetPrivateRepos() ([]Repo, error) {
	return nil, nil
}

func getGitHubClient(c *Config) *GithubClient {
	return &GithubClient{
		client:   github.NewClient(nil),
		username: c.Github.Username,
	}
}
