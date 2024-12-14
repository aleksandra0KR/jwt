package implementation

import "jwt/internal/repository"

type RefreshTokenUsecaseImplementation struct {
	repository repository.RefreshTokenRepository
}

func NewRefreshTokenUsecaseImplementation(repository repository.RefreshTokenRepository) *RefreshTokenUsecaseImplementation {
	return &RefreshTokenUsecaseImplementation{repository: repository}
}
