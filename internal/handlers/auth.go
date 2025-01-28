package handlers

import (
	"github.com/Vladislav557/auth/internal/models/http/request"
	"github.com/Vladislav557/auth/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type AuthHandler struct {
	rs service.RegistrationService
}

func (ah *AuthHandler) Register(ctx *gin.Context) {
	var req request.RegistrationRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		zap.L().Error("failed to parse registration request", zap.Error(err))
		return
	}
	response, err := ah.rs.Register(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		zap.L().Error("failed to parse registration request", zap.Error(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"AccessToken": response.AccessToken, "RefreshToken": response.RefreshToken})
}

//func (ah *AuthHandler) Login(ctx *gin.Context) {
//
//}
