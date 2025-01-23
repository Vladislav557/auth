package handlers

import (
	"github.com/Vladislav557/auth/internal/models/http/request"
	"github.com/Vladislav557/auth/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type RegistrationHandler struct {
	rs service.RegistrationService
}

func (rh *RegistrationHandler) Register(ctx *gin.Context) {
	var req request.RegistrationRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		zap.L().Error("failed to parse registration request", zap.Error(err))
		return
	}
	token, err := rh.rs.Register(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		zap.L().Error("failed to parse registration request", zap.Error(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
