package implementation

import (
	"jwt/domain"
	"jwt/internal/repository"
)

type UserUsecaseImplementation struct {
	repository repository.UserRepository
}

func NewUserUsecaseImplementation(repository repository.UserRepository) *UserUsecaseImplementation {
	return &UserUsecaseImplementation{repository: repository}
}

func (userUsecase *UserUsecaseImplementation) RegisterUser(user *domain.User) error {
	return nil
}

func (userUsecase *UserUsecaseImplementation) Auth(user *domain.User) (error, *domain.User) {
	return nil, nil
}
