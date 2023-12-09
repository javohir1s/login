package handler

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/asadbekGo/market_system/config"
	"github.com/asadbekGo/market_system/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	var loginRequest models.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		handleResponse(c, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	user, err := h.strg.User().GetByLoginAndPassword(context.Background(), &models.LoginRequest{
		Login:    loginRequest.Login,
		Password: loginRequest.Password,
	})
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			handleResponse(c, http.StatusUnauthorized, "Invalid login or password")
			return
		}

		log.Println("Error retrieving user:", err)
		handleResponse(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	if loginRequest.Password != user.Password {
		handleResponse(c, http.StatusUnauthorized, "Invalid login or password")
		return
	}

	resp := map[string]interface{}{
		"user":    user,
		"api_key": config.SecureApiKey,
	}

	handleResponse(c, http.StatusOK, resp)
}

func (h *Handler) RefreshToken(c *gin.Context) {
	var refreshToken models.UserPrimaryKey

	err := c.ShouldBindJSON(&refreshToken)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	err = h.strg.User().Refresh(ctx, refreshToken.Login)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusAccepted, "Token refreshed successfully")
}
