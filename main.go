package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/200", status200Handler)
	http.HandleFunc("/404", status404Handler)
	http.HandleFunc("/500", status500Handler)
	http.HandleFunc("/authenticated", authHandler)
	http.HandleFunc("/limited", limitHandler)

	fmt.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
