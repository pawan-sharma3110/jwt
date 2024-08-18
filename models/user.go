package models

import (
	"errors"
	"fmt"
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

func (u User) Validation() (uuid.UUID, error) {
	DB, _ := database.DbIn()
	var pass string
	query := `SELECT id,password FROM users WHERE email=$1 `
	err := DB.QueryRow(query, u.Email).Scan(&u.ID, &pass)
	if err != nil {
		return uuid.Nil, err
	}
	isValid := bcrypt.CompareHashAndPassword([]byte(pass), []byte(u.Password))
	if isValid != nil {
		return uuid.Nil, errors.New("invalid password")
	}
	return u.ID, nil
}
func AllUserGet() ([]User, error) {
	DB, _ := database.DbIn()
	query := `SELECT id,email FROM users`
	row, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	var users []User
	for row.Next() {
		var user User
		err = row.Scan(&user.ID, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := row.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}
	return users, nil

}
