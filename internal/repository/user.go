package repository

import (
	"jwt/domain"
)

type UserRepository interface {
	Register(user *domain.User) error
	Auth(user *domain.User) (error, *domain.User)
}
