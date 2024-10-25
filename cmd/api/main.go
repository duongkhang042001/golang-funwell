package main

import (
	"core/internal"
	"log"
)

func main() {
	app, err := internal.InitializeApplication("./config/config-local")
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
