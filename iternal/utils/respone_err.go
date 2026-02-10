package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ErrorBadRequest(w http.ResponseWriter, err error, str string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	x := fmt.Sprintf("%s : %s", str, err)
	json.NewEncoder(w).Encode(x)
}

func ErrorInternalServer(w http.ResponseWriter, err error, str string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	x := fmt.Sprintf("%s : %s", str, err)
	json.NewEncoder(w).Encode(x)
}

func ErrorMethodNotAllowed(w http.ResponseWriter, err error, str string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	x := fmt.Sprintf("%s : %s", str, err)
	json.NewEncoder(w).Encode(x)
}
