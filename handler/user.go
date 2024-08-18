package handler

import (
	"encoding/json"
	"fmt"
	"jwt/models"
	"net/http"
)

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "only get request allowed", http.StatusBadRequest)
		return
	}
	if r.URL.Path != "/all/users" {
		http.Error(w, "Invalid url", http.StatusBadRequest)
		return
	}
	users, err := models.AllUserGet()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, fmt.Sprintf("error:%v", err), http.StatusInternalServerError)
		return
	}

}
func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "only post request allowed", http.StatusBadRequest)
		return
	}
	if r.URL.Path != "/register" {
		http.Error(w, "Invalid url", http.StatusBadRequest)
		return
	}
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}
	id, err := user.SaveUser()
	if err != nil {
		http.Error(w, fmt.Sprintf("error:%v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "aplication/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(fmt.Sprintf("user register id:%v", id))
	if err != nil {
		http.Error(w, fmt.Sprintf("error:%v", err), http.StatusInternalServerError)
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "only post request allowed", http.StatusBadRequest)
		return
	}
	if r.URL.Path != "/login" {
		http.Error(w, "Invalid url", http.StatusBadRequest)
		return
	}
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}
	id, err := user.Validation()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "aplication/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(fmt.Sprintf("user login successfuly id:%v", id))
	if err != nil {
		http.Error(w, fmt.Sprintf("error:%v", err), http.StatusInternalServerError)
		return
	}
}
