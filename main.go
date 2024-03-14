package main

import (
	"fmt"

	"github.com/baileyjm02/muzz-backend-technical-exercise/api"
	"github.com/baileyjm02/muzz-backend-technical-exercise/match"
	"github.com/baileyjm02/muzz-backend-technical-exercise/server"
	"github.com/baileyjm02/muzz-backend-technical-exercise/swipe"
	"github.com/baileyjm02/muzz-backend-technical-exercise/user"
	"github.com/baileyjm02/muzz-backend-technical-exercise/utils"
)

func main() {
	fmt.Println("Hello, World!")

	user.AutoMigrate()
	swipe.AutoMigrate()
	match.AutoMigrate()

	go server.Start()
	go api.Start()

	stateCheck := make(chan struct{})
	channels := []chan struct{}{
		stateCheck,
	}

	utils.Preserve(channels)
}
