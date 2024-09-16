package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "<!DOCTYPE html><html><body>%s</body></html>", html.EscapeString(string(body)))

	} else {
		fmt.Fprintf(w, "<!DOCTYPE html><html><body><h1>Welcome to the home page!</h1></body></html>")
	}

}
