package main

import (
	"jwt/database"
	"jwt/handler"
	"log"
	"net/http"
)

func main() {
	DB, _ := database.DbIn()
	defer DB.Close()
	http.HandleFunc("/register", handler.Register)
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/all/users", handler.GetAllUser)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
