package handler

import (
	"github.com/gin-gonic/gin"
	"jwt/domain"
	"jwt/internal/usecase"
	"net/http"
)

type Handler struct {
	service *usecase.UseCase
}

func NewHandler(service *usecase.UseCase) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Handle() http.Handler {
	router := gin.Default()

	router.POST("/api/refreshToken/:guid", h.RefreshToken)
	router.POST("/api/auth/:guid", h.Auth)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(domain.NotImplementedStatusResponse,
			gin.H{"code": domain.NotImplementedStatusResponse, "error": "not implemented"})
	})
	return router
}
