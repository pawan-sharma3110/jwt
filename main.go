package main

import (
	"jwt/database"
	"jwt/handler"
	"log"
	"net/http"
)

func main() {
	database.DbIn()
	http.HandleFunc("/", handler.Register)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
