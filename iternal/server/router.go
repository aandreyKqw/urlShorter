package server

import "net/http"

func helloworld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func MakeNewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/helloworld", helloworld)

	return mux
}
