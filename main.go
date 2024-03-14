package main

import (
	"github.com/baileyjm02/muzz-backend-technical-exercise/match"
	"github.com/baileyjm02/muzz-backend-technical-exercise/server"
	"github.com/baileyjm02/muzz-backend-technical-exercise/swipe"
	"github.com/baileyjm02/muzz-backend-technical-exercise/user"
	"github.com/baileyjm02/muzz-backend-technical-exercise/utils"
)

// Main entry point for the application
func main() {
	// Ensure all tables are created / migrated
	user.AutoMigrate()
	swipe.AutoMigrate()
	match.AutoMigrate()

	// Start the API server
	go server.Start()

	// Example channel if we needed to block a graceful shutdown
	// process elsewhere
	exampleChannel := make(chan struct{})
	channels := []chan struct{}{
		exampleChannel,
	}

	utils.Preserve(channels)
}
