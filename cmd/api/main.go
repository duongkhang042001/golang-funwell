package main

import (
	"core/internal"
	"log"
)

func main() {
	cfgPath := "./config/config-local"

	app, err := internal.InitializeApplication(cfgPath)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
