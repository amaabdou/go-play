package main

import (
	"github.com/amaabdou/go-play-filesystem-http-browser/server"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("USAGE: go run test.go PORT")
	}

	port, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		log.Fatal("USAGE: go run test.go PORT")
	}

	server.Listen(port)
}
