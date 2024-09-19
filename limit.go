package main

import (
	"fmt"
	"net/http"

	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(rate.Limit(100), 30)

func limitHandler(w http.ResponseWriter, r *http.Request) {
	if !limiter.Allow() {
		http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
		return
	}
	fmt.Fprint(w, "Request accepted")
}
