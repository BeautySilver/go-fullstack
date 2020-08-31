package routes

import (
	"../../middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Uri     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

func Load() [] Route{
  routes:= userRoutes
  routes = append(routes, postsRoutes...)
	routes = append(routes, loginRoutes...)
  return routes
}

func SetupRoutes(r *mux.Router) *mux.Router {
	for _, route := range Load(){
		r.HandleFunc(route.Uri, route.Handler).Methods(route.Method)
	}
	return r
}

func SetupRoutesWithMiddlewares(r *mux.Router) *mux.Router {
	for _, route := range Load(){
		if route.AuthRequired{
			r.HandleFunc(route.Uri,
				middlewares.SetMiddlewareLogger(
					middlewares.SetMiddlewareJSON(
						middlewares.SetNiddlewareAuth(route.Handler))),

			).Methods(route.Method)
		}else{
			r.HandleFunc(route.Uri,
				middlewares.SetMiddlewareLogger(
					middlewares.SetMiddlewareJSON(route.Handler)),

			).Methods(route.Method)
		}
	}
	return r
}