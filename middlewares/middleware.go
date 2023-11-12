package api

import (
	"net/http"
)

func ExampleMiddleware(next http.Handler) http.Handler {
	return http.HandleFunc(func(w http.ResponseWriter , r *http.Request){

		next.ServerdHTTP(w ,r)
	})
}