package routes

import "github.com/gorilla/mux"

func IndexRoute(r *mux.Router) {
	api := r.PathPrefix("/api").Subrouter()

	AuthorRoutes(api)
}