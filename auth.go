package main

import (
	"fmt"
	"net/http"
	"os"
)

func authHandler(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !ok || !checkCredentials(username, password) {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	fmt.Fprintf(w, "Welcome, %s!", username)
}

func checkCredentials(username, password string) bool {
	return username == os.Getenv("AUTH_USERNAME") && password == os.Getenv("AUTH_PASSWORD")
}
