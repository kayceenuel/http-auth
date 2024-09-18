package main

import (
	"fmt"
	"net/http"
	"os"
)

func authHandler(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !ok || !checkCredentials(username, password) {
		w.Header().Set("WWW-Authenticate", `http-StatusUnauthorized`)
		return
	}
	fmt.Fprintf(w, "Hello, %s!", username)
}

func checkCredentials(username, password string) bool {
	return username == os.Getenv("AUTH_USERNAME") && password == os.Getenv("Auth_PASSWORD")
}
