package middlewares

import (
	"../auth"
	"../responses"
	"log"
	"net/http"
)

func SetMiddlewareLogger(next http.HandlerFunc)  http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\r%s %s%s %s", r.Method, r.Host, r.RequestURI, r.Proto)
		next(w, r)
	}
}

func SetMiddlewareJSON(next http.HandlerFunc)  http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			log.Println(err)
			responses.ERROR(w, http.StatusUnauthorized, err)
		}
		next(w, r)
	}
}