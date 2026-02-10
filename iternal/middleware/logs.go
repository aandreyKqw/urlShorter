package middleware

import (
	"log"
	"net/http"
)

func LogsMidllware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Host, r.Method, r.URL.String())
		next.ServeHTTP(w, r)
		log.Println("Sucsesful")
	})
}
