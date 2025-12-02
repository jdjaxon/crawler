package main

import (
	"log"
	"os"
)

func main() {
	argCount := len(os.Args)
	if argCount > 2 {
		log.Println("too many arguments provided")
		os.Exit(1)
	} else if argCount < 2 {
		log.Println("no website provided")
		os.Exit(1)
	}

	baseURL := os.Args[1]
	log.Printf("starting crawl of: %s\n", baseURL)
}
