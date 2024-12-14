package usecase

import (
	"jwt/internal/repository"
	"jwt/internal/usecase/implementation"
)

type UseCase struct {
	RefreshTokenUsecase
	UserUsecase
}

func NewUseCase(repository *repository.Repository) *UseCase {
	return &UseCase{
		RefreshTokenUsecase: implementation.NewRefreshTokenUsecaseImplementation(repository.RefreshTokenRepository),
		UserUsecase:         implementation.NewUserUsecaseImplementation(repository.UserRepository),
	}
}
