package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"../../middlewares"
)

type Route struct {
	Uri string
	Method string
	Handler func(w http.ResponseWriter, r *http.Request)
	AuthRequired bool
}

func Load() []Route{
	routes := usersRoutes
	routes = append(routes, postsRoutes...)
	routes = append(routes, loginRoutes...)
	return routes
}

func SetupRoutes(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		r.HandleFunc(route.Uri, route.Handler).Methods(route.Method)
	}
	return r
}

func SetupRoutesWithMiddlewares(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		if route.AuthRequired {
			r.HandleFunc(
				route.Uri,
				middlewares.SetMiddlewareLogger(
					middlewares.SetMiddlewareJSON(
						middlewares.SetMiddlewareAuthentication(route.Handler),
					),
				),
			).Methods(route.Method)
		} else {

			r.HandleFunc(
				route.Uri,
				middlewares.SetMiddlewareLogger(
					middlewares.SetMiddlewareJSON(
						route.Handler,
					),
				),
			).Methods(route.Method)
		}
	}
	return r
}