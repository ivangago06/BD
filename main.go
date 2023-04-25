package main

import (
	"BD/routes"
	"log"
	"net/http"
)

func main() {
	r := routes.Router()
	log.Println("Servidor Iniciado")
	log.Fatal(http.ListenAndServe(":8080", r))

}
