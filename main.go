package main

import (
	"fmt"

	"github.com/bah2830/brentahughes.com/github"
	"github.com/bah2830/brentahughes.com/webserver"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("Starting...")

	// Setup configuration
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	fmt.Println("Configuration loaded")

	/*
	   Start the Github poller. This is not required but will
	   asynchronously get Github stats without
	   wasting the users time during page load.
	*/
	github.StartPoller()

	// Start the web server
	webserver.Start()
}
