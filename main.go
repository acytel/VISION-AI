package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
)

func main() {
	// Create a new PocketBase app instance.
	app := pocketbase.New()

	// Start the PocketBase framework.
	// This is the core command that gets the server running.
	// We have removed the extra logging hooks for maximum compatibility.
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

