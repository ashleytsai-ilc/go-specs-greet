package httpserver

import (
	"fmt"
	"go-specs-greet/domain/interactions"
	"net/http"
)

func Hanler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, interactions.Greet(name))
}
