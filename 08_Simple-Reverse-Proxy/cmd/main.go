package main

import (
	"log"

	"github.com/Shubham-Kumar1/08_Simple-Reverse-Proxy/internal/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatalf("could not start the server: %v", err)
	}
}
