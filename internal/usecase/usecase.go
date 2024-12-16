package usecase

import (
	"jwt/domain"
	"jwt/internal/repository"
	"jwt/internal/usecase/implementation"
)

type UseCase interface {
	RefreshToken(*domain.User) (error, *domain.User)
	Auth(*domain.User) error
}

func NewUseCase(repository *repository.Repository) UseCase {
	return implementation.NewUsecaseImplementation(repository)
}
