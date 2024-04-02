package httpserver

import (
	"fmt"
	"go-specs-greet/domain/interactions"
	"net/http"
)

func Handler() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/greet", replyWith(interactions.Greet))
	router.HandleFunc("/curse", replyWith(interactions.Curse))
	return router
}

func replyWith(f func(name string) (interaction string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		fmt.Fprint(w, f(name))
	}
}
