package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Serv struct {
	serv *http.Server
}

func MakeNewServ(addr string, handler http.Handler) *Serv {
	return &Serv{
		serv: &http.Server{
			Addr:    addr,
			Handler: handler,
		},
	}
}

func (server *Serv) Run() {
	log.Println("Запустили сервер на порту", server.serv.Addr)
	go func() {
		if err := server.serv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
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

	if err := server.serv.Shutdown(ctx); err != nil {
		log.Fatal("Ошибка при выключении сервера", err)
	}
	log.Println("Успешно завершили все операции и отключились")

}
