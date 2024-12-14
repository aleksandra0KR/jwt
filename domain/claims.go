package domain

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Guid  string `json:"sub"`
	IP    string `json:"ip"`
	Email string `json:"email"`
	jwt.StandardClaims
}
