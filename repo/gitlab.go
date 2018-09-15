package repo

import (
	"strconv"

	"github.com/xanzy/go-gitlab"
)

type GitlabClient struct {
	client   *gitlab.Client
	username string
	userID   string
	user     *gitlab.User
}

func (c *GitlabClient) GetPublicRepos() ([]*Repo, error) {
	opt := &gitlab.ListProjectsOptions{
		Visibility: gitlab.Visibility(gitlab.PublicVisibility),
		Simple:     gitlab.Bool(true),
	}

	r, _, err := c.client.Projects.ListUserProjects(c.userID, opt)
	if err != nil {
		return nil, err
	}

	repos := make([]*Repo, 0)
	for _, repo := range r {
		repos = append(repos, &Repo{
			Source:       "gitlab",
			Name:         repo.Name,
			URL:          repo.WebURL,
			Contribution: (repo.ForkedFromProject != nil),
		})
	}

	return repos, nil
}

func (c *GitlabClient) GetContributions() ([]*Repo, error) {
	return nil, nil
}

func (c *GitlabClient) getUser() (*gitlab.User, error) {
	id, _ := strconv.Atoi(c.userID)
	user, _, err := c.client.Users.GetUser(id)
	return user, err
}

func getGitlabClient(c *Config) *GitlabClient {
	return &GitlabClient{
		client:   gitlab.NewClient(nil, ""),
		username: c.Gitlab.Username,
		userID:   c.Gitlab.Token,
	}
}
