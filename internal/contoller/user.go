package handler

import (
	"github.com/gin-gonic/gin"
	"jwt/domain"
	"time"
)

func (h *Handler) Auth(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(domain.BadRequestStatusResponse,
			gin.H{"error": gin.H{"code": domain.BadRequestStatusResponse, "text": "invalid input"}})
		return
	}
	guid := c.Param("guid")
	user.Guid = guid

	err := (*h.service).Auth(&user)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(domain.UnauthorizedStatusResponse,
				gin.H{"error": gin.H{"code": domain.UnauthorizedStatusResponse, "text": err.Error()}})
			return
		}
		c.JSON(domain.InternalServerErrorStatusResponse,
			gin.H{"error": gin.H{"code": domain.InternalServerErrorStatusResponse, "text": err.Error()}})
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	ex := time.Now().Add(1 * time.Hour)
	c.SetCookie("accessToken", user.AccessToken, int(expirationTime.Unix()), "/", "localhost", false, true)
	c.SetCookie("refreshToken", user.RefreshToken, int(ex.Unix()), "/", "localhost", false, true)

	c.JSON(domain.SuccessfulStatusResponse,
		gin.H{"response": gin.H{"accessToken": user.AccessToken, "refreshToken": user.RefreshToken}})
}

func (h *Handler) RefreshToken(c *gin.Context) {
	accessToken, err := c.Cookie("accessToken")
	if err != nil {
		c.JSON(domain.UnauthorizedStatusResponse,
			gin.H{"error": gin.H{"code": domain.UnauthorizedStatusResponse, "text": "unauthorized"}})
		return
	}

	refreshToken, err := c.Cookie("refreshToken")
	if err != nil {
		c.JSON(domain.UnauthorizedStatusResponse,
			gin.H{"error": gin.H{"code": domain.UnauthorizedStatusResponse, "text": "unauthorized"}})
		return
	}

	var user *domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(domain.BadRequestStatusResponse,
			gin.H{"error": gin.H{"code": domain.BadRequestStatusResponse, "text": "invalid input"}})
		return
	}
	guid := c.Param("guid")
	user.Guid = guid

	user.AccessToken = accessToken
	user.RefreshToken = refreshToken
	err, user = (*h.service).RefreshToken(user)
	if err != nil {
		c.JSON(domain.BadRequestStatusResponse,
			gin.H{"error": gin.H{"code": domain.BadRequestStatusResponse, "text": "invalid input"}})
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	ex := time.Now().Add(1 * time.Hour)
	c.SetCookie("accessToken", user.AccessToken, int(expirationTime.Unix()), "/", "localhost", false, true)
	c.SetCookie("refreshToken", user.RefreshToken, int(ex.Unix()), "/", "localhost", false, true)

	c.JSON(domain.SuccessfulStatusResponse,
		gin.H{"response": gin.H{"accessToken": user.AccessToken, "refreshToken": user.RefreshToken}})
}
