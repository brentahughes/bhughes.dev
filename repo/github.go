package repo

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/github"
)

type GithubClient struct {
	client   *github.Client
	username string
}

func (c *GithubClient) GetPublicRepos() ([]*Repo, error) {
	r, _, err := c.client.Repositories.List(context.Background(), c.username, nil)
	if err != nil {
		return nil, err
	}

	repos := make([]*Repo, 0)
	for _, repo := range r {
		if *repo.Fork {
			continue
		}

		repos = append(repos, &Repo{
			Source: "github",
			Name:   *repo.Name,
			URL:    *repo.SVNURL,
		})
	}

	return repos, nil
}

func (c *GithubClient) GetContributions() ([]*Repo, error) {
	r, _, err := c.client.Search.Issues(
		context.Background(),
		fmt.Sprintf("author:%s is:pr", c.username),
		nil,
	)
	if err != nil {
		return nil, err
	}

	reposURLs := make(map[string]bool, 0)
	for _, issue := range r.Issues {
		if _, ok := reposURLs[*issue.RepositoryURL]; ok {
			continue
		}

		reposURLs[*issue.RepositoryURL] = true
	}

	repos := make([]*Repo, 0)
	for repoURL := range reposURLs {
		repoPath := strings.TrimPrefix(repoURL, "https://api.github.com/repos/")
		parts := strings.Split(repoPath, "/")
		repos = append(repos, &Repo{
			Source:       "github",
			Name:         parts[1],
			URL:          fmt.Sprintf("https://github.com/%s/%s", parts[0], parts[1]),
			Contribution: true,
		})
	}
	return repos, nil
}

func getGitHubClient(c *Config) *GithubClient {
	client := github.NewClient(nil)

	return &GithubClient{
		client:   client,
		username: c.Github.Username,
	}
}
