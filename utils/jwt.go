package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

const secrectKey string = "jwt_token"

func GernateJwt(id uuid.UUID, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         id,
		"email":      email,
		"expireTime": time.Now().Add(time.Minute * 5).Unix(),
	})

	return token.SignedString([]byte(secrectKey))
}
func VerifyJwt(token string) (uuid.UUID, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return uuid.Nil, errors.New("unexcpted singin method")
		}
		return []byte(secrectKey), nil
	})
	if err != nil {
		return uuid.Nil, err
	}
	tokenValid := parsedToken.Valid
	if !tokenValid {
		return uuid.Nil, errors.New("invalid token")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.Nil, errors.New("invalid token claims")
	}
	// email := claims["email"].(string)
	id := claims["id"].(uuid.UUID)
	return id, nil
}
