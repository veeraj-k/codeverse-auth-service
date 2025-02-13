package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {

	userRepo := NewUserRepository(db)
	authService := NewAuthService(userRepo)
	authHandler := NewAuthHandler(authService)

	r.POST("/api/register", authHandler.Register)
	r.POST("/api/login", authHandler.Login)

}
