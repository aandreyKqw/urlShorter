package handlers

import (
	"errors"
	"net/http"
	"urlshort/iternal/utils"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func checkShortInDB(url string, conn *pgxpool.Pool, r *http.Request) (string, error) {
	var value string
	err := conn.QueryRow(r.Context(), "SELECT original_url FROM urls WHERE short_code=$1", url).Scan(&value)
	if err == nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return value, nil
		}
	}
	err = errors.New("Такой строки нет")
	return "", err
}

func SwapOnRealSite(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		url := r.URL.String()
		url = url[1:]

		longUrl, checkBool := checkShortInDB(url, conn, r)
		if checkBool != nil {
			w.Header().Set("Content-Type", "application/json")
			utils.ErrorBadRequest(w, checkBool, "Видмо не верный адрес")
			return
		}
		w.Header().Set("Location", longUrl)
		w.WriteHeader(http.StatusFound)
	default:
		err := errors.New("Не тот метод")
		utils.ErrorMethodNotAllowed(w, err, "Нужен только GET")
	}

}
