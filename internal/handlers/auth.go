package handlers

import (
	"github.com/Vladislav557/auth/internal/lib/jwt"
	"github.com/Vladislav557/auth/internal/models/http/request"
	"github.com/Vladislav557/auth/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type AuthHandler struct {
	authorizationService service.AuthorizationService
}

func (authHandler *AuthHandler) SingUp(ctx *gin.Context) {
	var req request.SingUpRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		zap.L().Error("failed to parse registration request", zap.Error(err))
		return
	}
	err := authHandler.authorizationService.Register(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		zap.L().Error("failed to parse registration request", zap.Error(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "accept your registration by email"})
}

func (authHandler *AuthHandler) SingIn(ctx *gin.Context) {
	var req request.SingInRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		zap.L().Error("failed to parse registration request", zap.Error(err))
		return
	}
	res, err := authHandler.authorizationService.LoginByEmail(req.Email, req.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		zap.L().Error("failed to parse registration request", zap.Error(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (authHandler *AuthHandler) RefreshTokens(ctx *gin.Context) {
	//TODO
}

func (authHandler *AuthHandler) Logout(ctx *gin.Context) {
	authorization := ctx.GetHeader("Authorization")
	claims, err := jwt.ParseToken(authorization)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		zap.L().Error("failed to parse registration request", zap.Error(err))
		return
	}
	err = authHandler.authorizationService.Logout(ctx, claims)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		zap.L().Error("failed to parse registration request", zap.Error(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": true})
}

func (authHandler *AuthHandler) ConfirmEmail(ctx *gin.Context) {
	uuid := ctx.Query("user")
	err := authHandler.authorizationService.Confirm(ctx, uuid)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		zap.L().Error("failed to confirm email", zap.Error(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "email confirmed"})
}
