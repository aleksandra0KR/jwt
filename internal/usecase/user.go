package usecase

import "jwt/domain"

type UserUsecase interface {
	RegisterUser(user *domain.User) error
	Auth(user *domain.User) (error, *domain.User)
}
