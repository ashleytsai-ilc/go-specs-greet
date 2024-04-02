package main

import (
	"go-specs-greet/adapters/httpserver"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(httpserver.Hanler)
	http.ListenAndServe(":8080", handler)
}
