package main

import (
	"log"
	"data-preparer/internal/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal("Failed to create app:", err)
	}

	log.Println("DataPreparer started")
	if err := a.Run(); err != nil {
		log.Fatal("App run error:", err)
	}
}