package main

import (
	"log"

	"github.com/morfo-si/beam/internal/server"
)

// Main function to set up the GoFiber server
func main() {
	srv := server.NewACEServer()

	// Start the server on port 8081
	log.Fatal(srv.Start())
}
