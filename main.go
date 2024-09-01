package main

import (
	"log"

	"github.com/dario-labs/srv/internal/adapters"
)

func main() {
	server := adapters.NewServer()
	if err := server.Run(":3000"); err != nil {
		log.Fatalln("Can't start the server")
	}
}
