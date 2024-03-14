package utils

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Preserve is a function that listens for exit signals and closes channels.
// This allows us to gracefully close the application, within Docker we normally have 10
// seconds to close the application before it is forcefully closed. Therefore we can use
// 'channels' to gracefully close areas of the application. E.g. the API server and DB.
func Preserve(channels []chan struct{}) {
	// Create a wait group so we can cleanly block the thread while code remains readable.
	waiter := new(sync.WaitGroup)
	waiter.Add(1)

	// Create a channel that listens for exit signals
	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)

	go func() {
		// Block the thread until we receive a signal in the channel
		<-exitSignal

		fmt.Println("Cleaning up and closing")

		// Nicely close channels. This allows us to gracefully close the application areas,
		// else where by using an async within very similar to this one.
		for _, ch := range channels {
			close(ch)
		}

		waiter.Done()
	}()

	// Wait until the wait group is done.
	waiter.Wait()

	// Exit the application
	os.Exit(0)
}
