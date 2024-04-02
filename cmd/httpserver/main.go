package main

import (
	"go-specs-greet/adapters/httpserver"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", httpserver.Handler())
}
