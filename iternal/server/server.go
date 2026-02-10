package Server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"urlshort/iternal/middleware"
)

func MakeNewServ(addr string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: middleware.LogsMidllware(handler),
	}
}

func Run(server *http.Server) {
	log.Println("Запустили сервер на порту", server.Addr)
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln("Ошибка при старте сервера", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("Начинаем процесс выключения сервера")
	log.Println("Ожидаем окончания последних операций")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Ошибка при выключении сервера", err)
	}
	log.Println("Успешно завершили все операции и отключились")

}
