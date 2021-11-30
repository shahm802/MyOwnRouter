package main

import (
	"net/http"

	"github.com/shahm802/GoRoute/router"
)

func main() {
	r := &router.Router{}
	r.Route("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("The best Router!"))
	})
	http.ListenAndServe(":8080", r)
}
