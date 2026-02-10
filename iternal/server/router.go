package Server

import (
	"net/http"
	"urlshort/iternal/handlers"

	"github.com/jackc/pgx/v5/pgxpool"
)

func MakeNewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	return mux
}

func ListHandlers(conn *pgxpool.Pool, mux *http.ServeMux) {
	mux.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) { handlers.Create(conn, w, r) })
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { handlers.SwapOnRealSite(conn, w, r) })
}
