package main

import (
	"context"
	"log"
)

var ctx = context.Background()

func main() {
	app, err := NewApp()

	if err != nil {
		log.Fatal(err)
		return
	}

	defer app.Stop()

	app.Run()
}
