package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"urlshort/iternal/db"
	Server "urlshort/iternal/server"
)

func main() {
	f, err := os.OpenFile("../logs.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Что-то не так с файлом для логов", err)
	}
	defer f.Close()
	originalWriter := log.Writer()
	multiWriter := io.MultiWriter(originalWriter, f)
	log.SetOutput(multiWriter)

	conn := db.Connect()

	mux := Server.MakeNewRouter()
	server := Server.MakeNewServ(":8080", mux)
	Server.ListHandlers(conn, mux)

	Server.Run(server)

}
