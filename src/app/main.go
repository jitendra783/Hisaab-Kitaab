package main

import (
	"flag"
	"hisaab-kitaab/api"
	"hisaab-kitaab/pkg/config"
	"log"
	"os"
)

func main() {

	// Set the environment, by default development
	flag.Usage = func() {
		log.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	environment := "local"
	host := os.Getenv("SERVER_HOST")
	if host != "" {
		environment = "server"
	}

	config.LoadConfig(environment)
	if err := api.Start(); err != nil {
		log.Fatal("Failed to start server, err:", err)
		os.Exit(1)
	}
}
