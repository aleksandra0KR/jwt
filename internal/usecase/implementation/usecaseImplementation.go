package implementation

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"jwt/domain"
	"jwt/internal/contoller/middleware"
	"jwt/internal/repository"
	"net/smtp"
	"os"
)

type UsecaseImplementation struct {
	repository *repository.Repository
}

func NewUsecaseImplementation(repository *repository.Repository) *UsecaseImplementation {
	return &UsecaseImplementation{repository: repository}
}

func (usecase *UsecaseImplementation) Auth(user *domain.User) error {
	var jwt middleware.JWT
	accessToken, err := jwt.GenerateToken(user)
	if err != nil {
		log.Error(err)
		return err
	}
	user.AccessToken = accessToken

	refreshToken, err := jwt.GenerateToken(user)
	if err != nil {
		log.Error(err)
		return err
	}
	user.RefreshToken = refreshToken

	token := domain.RefreshToken{
		RefreshToken: refreshToken,
		Guid:         user.Guid,
	}

	err = (*usecase.repository).Auth(&token)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (usecase *UsecaseImplementation) RefreshToken(user *domain.User) (error, *domain.User) {
	_, err := middleware.JWT{}.ParseToken(user.AccessToken)
	if err != nil {
		return err, nil
	}

	var refreshToken *domain.RefreshToken
	refreshToken, err = (*usecase.repository).GetRefreshToken(user.Guid)
	if err != nil {
		return err, nil
	}

	if refreshToken.RefreshToken != user.RefreshToken {
		return fmt.Errorf("wrong refresh token"), nil
	}

	userFromDB, err := (*usecase.repository).GetUserByGuid(user.Guid)
	if err != nil {
		return err, nil
	}

	if userFromDB.IP != user.IP {
		err = usecase.SendEmail(user)
		if err != nil {
			return err, nil
		}
	}

	err = (*usecase.repository).DeleteRefreshToken(user.Guid)
	if err != nil {
		return err, nil
	}

	err = usecase.Auth(user)
	if err != nil {
		return err, nil
	}
	return nil, user
}

func (usecase *UsecaseImplementation) SendEmail(user *domain.User) error {
	from := os.Getenv("EMAIL_FROM")
	password := os.Getenv("PASSWORD_EMAIL")

	to := []string{user.Email}
	host := "smtp.gmail.com"
	port := "587"
	subject := "Subject: login\n"
	body := "You have logged in from another device"
	message := []byte(subject + body)

	auth := smtp.PlainAuth("", from, password, host)
	err := smtp.SendMail(host+":"+port, auth, from, to, message)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
