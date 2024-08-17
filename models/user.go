package models

import (
	"jwt/database"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

func (u User) SaveUser() (id string, err error) {
	DB, _ := database.DbIn()
	password, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return "", err
	}
	u.Password = string(password)
	u.ID = uuid.New()
	query := `INSERT INTO users(id,email,password)VALUES($1,$2,$3)RETURNING id`
	stmt, err := DB.Prepare(query)
	if err != nil {
		return "", err
	}

	err = stmt.QueryRow(u.ID, u.Email, u.Password).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}
