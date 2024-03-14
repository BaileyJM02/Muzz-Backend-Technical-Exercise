package main

import (
	"fmt"

	"github.com/baileyjm02/muzz-backend-technical-exercise/api"
	"github.com/baileyjm02/muzz-backend-technical-exercise/utils"
)

func main() {
	fmt.Println("Hello, World!")

	api.Start()

	stateCheck := make(chan struct{})
	channels := []chan struct{}{
		stateCheck,
	}

	utils.Preserve(channels)
}
