package main

import (
	"jwt/database"
	"jwt/handler"
	"jwt/middleware"
	"log"
	"net/http"
)

func main() {
	DB, _ := database.DbIn()
	defer DB.Close()
	http.HandleFunc("/register", handler.Register)
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/all/users", middleware.Auth(handler.GetAllUser))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
