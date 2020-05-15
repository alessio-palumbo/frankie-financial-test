package main

import (
	"github.com/alessio-palumbo/frankie-financial-test/pkg/handlers"
)

// @title Frankie-Financial-Test
// @version 1.0
// @description This API allows developers to test the Universal SDK output to ensure it looks right.

func main() {
	// Create a gin router with default middleware and register routes.
	r := handlers.SetupRouter()

	r.Run()
}
