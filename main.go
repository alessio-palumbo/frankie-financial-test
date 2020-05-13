package main

import "github.com/alessio-palumbo/frankie-financial-test/handlers"

func main() {
	// Create a gin router with default middleware and register routes.
	r := handlers.SetupRouter()

	r.Run()
}
