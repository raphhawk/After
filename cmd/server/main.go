package main

import (
	"log"

	"github.com/raphhawk/proglog/internal/server"
)

func main() {
	srv := server.NewHttpServer(":8080")
	log.Println("Listening on port: 8080")
	log.Fatal(srv.ListenAndServe())
}
