package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect() *pgxpool.Pool {

	Conn, err := pgxpool.New(context.Background(), os.Getenv("DB_CONN"))
	if err != nil {
		log.Println("Ошибка подключения к БД", err)
		return nil
	}

	if err := Conn.Ping(context.Background()); err != nil {
		log.Println("БД не пингуется", err)
		return nil
	}

	time.Sleep(4 * time.Second)
	log.Println("Успешно подключились к БД")

	return Conn
}
