package main

import (
	"fmt"

	"github.com/baileyjm02/muzz-backend-technical-exercise/api"
	"github.com/baileyjm02/muzz-backend-technical-exercise/database"
	"github.com/baileyjm02/muzz-backend-technical-exercise/utils"
)

func main() {
	fmt.Println("Hello, World!")

	database.Start()

	api.Start()

	stateCheck := make(chan struct{})
	channels := []chan struct{}{
		stateCheck,
	}

	utils.Preserve(channels)
}
