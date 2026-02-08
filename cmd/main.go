package main

import (
	"urlshort/iternal/db"
	"urlshort/iternal/server"
)

func main() {
	_ = db.Connect()

	mux := server.MakeNewRouter()
	server := server.MakeNewServ(":8080", mux)

	server.Run()

}
