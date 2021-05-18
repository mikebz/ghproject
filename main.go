package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("starting ghproject")
	fmt.Println("GITHUB_TOKEN: ", os.Getenv("GITHUB_TOKEN"))
	os.Exit(0)
}
