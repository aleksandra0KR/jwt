package domain

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	IP    string `json:"ip"`
	Email string `json:"email"`
	jwt.StandardClaims
}
