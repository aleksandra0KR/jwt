package middleware

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"io"
	"jwt/domain"
	"os"
	"time"
)

type JWT struct{}

func (JWT) GenerateToken(user *domain.User) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &domain.Claims{
		IP:    user.IP,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			Subject:   user.Guid,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		log.Errorf("error generating token: %v", err)
	}

	return tokenString, err
}

func (JWT) GenerateRefreshToken(accessToken string) (string, error) {
	sha1 := sha1.New()
	_, err := io.WriteString(sha1, os.Getenv("SECRET_KEY"))
	if err != nil {
		return "", err
	}

	salt := string(sha1.Sum(nil))[0:16]
	block, err := aes.NewCipher([]byte(salt))
	if err != nil {
		fmt.Println(err.Error())

		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return "", err
	}

	refreshToken := base64.URLEncoding.EncodeToString(gcm.Seal(nonce, nonce, []byte(accessToken), nil))
	return refreshToken, nil
}

func (JWT) ParseToken(tokenString string) (claims *domain.Claims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &domain.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*domain.Claims)
	if !ok {
		return nil, err
	}

	return claims, nil
}
