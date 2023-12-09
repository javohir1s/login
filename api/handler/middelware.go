package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/asadbekGo/market_system/config"
	"github.com/gin-gonic/gin"
)

func checkApiKeyExpiration(c *gin.Context) bool {
	expireDuration := config.ApiKeyExpiredAt

	expireTime := time.Now().Add(expireDuration)

	if time.Now().After(expireTime) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Expired API key"})
		return false
	}

	return true
}

func (h *Handler) CheckPasswordMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !checkApiKeyExpiration(c) {
			return
		}

		password := c.GetHeader("API-KEY")
		if password != config.SecureApiKey {
			c.AbortWithError(http.StatusForbidden, errors.New("The request requires user authentication."))
			return
		}

		c.Next()
	}
}
