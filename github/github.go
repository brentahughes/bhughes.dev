package github

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/github"
	"github.com/spf13/viper"
)

var repoCache []Repo
var contributionCache []Repo

type Repo struct {
	Name string
	URL  string
}

func StartPoller() {
	fmt.Println("Starting github api poller...")

	// Run once every hour
	ticker := time.NewTicker(time.Hour * 1)
	go func() {
		updateRepos()
		for range ticker.C {
			updateRepos()
		}
	}()
}

func updateRepos() {
	fmt.Println("Updating repo cache")

	ctx := context.Background()
	client := github.NewClient(nil)
	repos, _, err := client.Repositories.List(ctx, viper.GetString("github.username"), nil)
	if err != nil {
		fmt.Printf("Error updating repo cache. %s", err)
	}

	repoCache = []Repo{}
	contributionCache = []Repo{}
	for _, repo := range repos {
		newRepo := Repo{
			Name: *repo.Name,
			URL:  *repo.HTMLURL,
		}

		if *repo.Fork == false {
			repoCache = append(repoCache, newRepo)
		} else {
			contributionCache = append(contributionCache, newRepo)
		}
	}
}

func GetRepos() []Repo {
	return repoCache
}

func GetContributions() []Repo {
	return contributionCache
}
