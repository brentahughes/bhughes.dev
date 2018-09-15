package main

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"gitlab.com/bah2830/brentahughes.com/repo"
	"gitlab.com/bah2830/brentahughes.com/webserver"
)

func main() {
	fmt.Println("Starting...")

	// Setup configuration
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	fmt.Println("Configuration loaded")

	// Start the Github poller. This is not required but will
	// asynchronously get Github stats without
	// wasting the users time during page load.
	repoClient := repo.GetClient(&repo.Config{
		Github: repo.RepoConfig{
			Username: viper.GetString("github.username"),
			Token:    viper.GetString("github.token"),
		},
		Gitlab: repo.RepoConfig{
			Username: viper.GetString("gitlab.username"),
			Token:    viper.GetString("gitlab.user_id"),
		},
	})

	// Make sure we can get the repos
	log.Println("Verifying remote repo connections are live")
	repos, err := repoClient.GetRepos(true)
	if err != nil {
		panic(err)
	}
	log.Printf("Found %d repos\n", len(repos))

	// Start the pller
	log.Println("Starting repo polling")
	go repoClient.Poll(6 * time.Hour)

	// Start the web server
	log.Println("Starting webserver")
	w := webserver.GetWebserver(repoClient)
	w.Start()
}
