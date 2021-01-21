package router

import "github.com/gorilla/mux"

func NEW() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	return r
}

