package main

import (
	"fmt"
	"html"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.Method == http.MethodPost {
		handlePost(w, r)
	} else {
		handleGet(w, r)
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
	fmt.Fprintf(w, "Query parameters:")
	for key, values := range r.URL.Query() {
		for _, value := range values {
			fmt.Fprintf(w, "%s: %s", html.EscapeString(key), html.EscapeString(value))
		}
	}
	fmt.Fprintf(w, "")
}
func handlePost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Received POST data")
	for key, values := range r.Form {
		for _, value := range values {
			fmt.Fprintf(w, "%s: %s", html.EscapeString(key), html.EscapeString(value))
		}
	}
	fmt.Fprintf(w, "")
}

func status200Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "200 OK")
}

func status404Handler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func status500Handler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}
