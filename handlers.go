package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
)

// Handler processes both GET and POST requests
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	// HTML response structure
	fmt.Fprintf(w, "<!DOCTYPE html><html><body>")
	fmt.Fprintf(w, "<h1>Welcome to the home page!</h1>")

	// Handle POST request
	if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}
		// Escape and print the posted content
		fmt.Fprintf(w, "<p>Posted content: %s</p>", html.EscapeString(string(body)))
	}

	// Display query parameters in GET requests
	fmt.Fprintf(w, "<h2>Query Parameters:</h2><ul>")
	for key, values := range r.URL.Query() {
		for _, value := range values {
			fmt.Fprintf(w, "<li>%s: %s</li>", html.EscapeString(key), html.EscapeString(value))
		}
	}
	fmt.Fprintf(w, "</ul></body></html>")
}
