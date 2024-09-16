package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/200", status200Handler)
	http.HandleFunc("/404", status404Handler)
	http.HandleFunc("/500", status500Handler)

	fmt.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
func status200Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "200 OK")
}
func status404Handler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
func status500Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "500 Internal Server Error")
}
