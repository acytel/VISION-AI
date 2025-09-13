package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	// This is a "lifecycle hook" that runs right before the web server starts.
	// We are using it to print a confirmation message in the logs.
	// The previous function `OnAfterBootstrap` was renamed to `OnBeforeServe`
	// in a recent PocketBase update.
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		log.Println("----------------------------------------")
		log.Println("  Project VISION AI Backend is starting...")
		// e.Server.URL gives us the actual address the server is running on.
		log.Println("  Admin UI will be available at:", e.Server.URL)
		log.Println("----------------------------------------")
		return nil
	})

	// Start the PocketBase framework.
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

