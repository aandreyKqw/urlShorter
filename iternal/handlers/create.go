package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
	"urlshort/iternal/models"
	"urlshort/iternal/utils"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func insertIntoTable(conn *pgxpool.Pool, long string, short string, r *http.Request) error {
	_, err := conn.Exec(r.Context(), "INSERT INTO urls (original_url,short_code) values ($1,$2);", long, short)
	if err != nil {
		return err
	}
	return nil
}

func makeShortUrl(url string) string {
	var url2 string
	urlslice := strings.Split(url, ".")
	urlslice = urlslice[1 : len(urlslice)-1]
	if len(urlslice) > 1 {
		for i := 0; i < len(urlslice); i++ {
			if i != len(urlslice)-1 {
				url2 += urlslice[i] + "."
			} else {
				url2 += urlslice[i]
			}
		}
	} else {
		url2 = urlslice[0]
	}

	sourse := rand.New(rand.NewSource(time.Now().UnixNano()))
	r := rand.New(sourse)

	return fmt.Sprintf("%s%s%s%d", string(url2[0]), string(url2[len(url2)-1]), string(url2[1]), r.Intn(100))
}

func checkLongInDB(url string, conn *pgxpool.Pool, r *http.Request) bool {

	err := conn.QueryRow(r.Context(), "SELECT * FROM urls WHERE original_url=$1", url).Scan()
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return true
		}
	}

	return false
}

func checkLongValid(url string) (bool, error) {
	h := "https://"

	if url[:8] != h {
		err := fmt.Errorf("Ссылка должна начинаться с %s", h)
		return false, err
	}

	return true, nil
}

func Create(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		ur := models.StructDecodRBody{}
		err := json.NewDecoder(r.Body).Decode(&ur)
		if err != nil {
			utils.ErrorBadRequest(w, err, "Ошибка раскодировки JSON")
			return
		}
		checkBool, err := checkLongValid(ur.Long)
		if checkBool == false {
			utils.ErrorBadRequest(w, err, "Адрес не валиден")
			return
		}

		checkBool = checkLongInDB(ur.Long, conn, r)
		if checkBool == true {
			w.Header().Set("content-type", "application/json")
			json.NewEncoder(w).Encode("Такая строка есть")
			return
		}

		shorturl := makeShortUrl(ur.Long)

		if err = insertIntoTable(conn, ur.Long, shorturl, r); err != nil {
			utils.ErrorInternalServer(w, err, "Ошибка сервера")
			return
		}

		fmt.Fprintf(w, "Добавлена в таблицу %s", ur.Long)
		fmt.Fprintf(w, "Короткая ссылка %s", shorturl)

	default:
		err := errors.New("Не тот метод")
		utils.ErrorMethodNotAllowed(w, err, "Нужен только POST")
	}

}
