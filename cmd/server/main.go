package main

import (
	"balance/internal/app/server"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
