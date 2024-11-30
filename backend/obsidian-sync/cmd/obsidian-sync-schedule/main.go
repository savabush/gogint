package main

import (
	"time"

	. "github.com/savabush/obsidian-sync/internal/app"
	. "github.com/savabush/obsidian-sync/internal/config"
)

// main is the entry point of the obsidian-sync scheduler application.
// It sets up a ticker to run the App() function at regular intervals
// defined by Settings.APP.SCHEDULE. The scheduler can be gracefully
// stopped by sending a signal to the quit channel.
func main() {
	ticker := time.NewTicker(time.Duration(Settings.APP.SCHEDULE) * time.Minute)
	quit := make(chan struct{})
	Logger.Infof("Starting obsidian-sync scheduler. Starts every %v minutes", Settings.APP.SCHEDULE)
	for {
		select {
		case <-ticker.C:
			App()
		case <-quit:
			Logger.Info("Stopping obsidian-sync scheduler")
			ticker.Stop()
			return
		}
	}
}
