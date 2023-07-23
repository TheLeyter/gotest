package main

import (
	"log"
)

func main() {
	app, err := NewApp()

	if err != nil {
		log.Fatal(err)
		return
	}

	defer app.Stop()

	app.Run()
}
