package router

import (
	"context"
	"net/http"
	"regexp"
)

type Router struct {
	Routes []RouteEntry
}

func (sr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//loop over all routes in router
	//if match - then great - use the handler func
	//else return not found
	for _, e := range sr.Routes {
		match := e.Match(r)
		if match == nil {
			continue
		}
		ctx := context.WithValue(r.Context(), "params", match)
		e.Handler.ServeHTTP(w, r.WithContext(ctx))
		return
	}

	http.NotFound(w, r)
}

func (rtr *Router) Route(method string, path string, handlerFunc http.HandlerFunc) {

	exactPath := regexp.MustCompile("^" + path + "$")

	currRoute := RouteEntry{
		Method:  method,
		Path:    exactPath,
		Handler: handlerFunc,
	}
	rtr.Routes = append(rtr.Routes, currRoute)
}
