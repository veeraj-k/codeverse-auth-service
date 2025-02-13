package admin

import (
	"auth-service/internal/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {

	userRepo := auth.NewUserRepository(db)
	adminAuthService := NewAdminAuthService(userRepo)
	adminHandler := NewAdminAuthHandler(adminAuthService)

	ar := r.Group("/api/admin", AdminAuthMiddleware([]string{"SUPER_ADMIN", "ADMIN"}))

	ar.POST("/roles", adminHandler.CreateRole)
	ar.GET("/roles", adminHandler.GetRoles)
	ar.DELETE("/roles/:name", adminHandler.DeleteRoleByName)
	ar.GET("/roles/:name", adminHandler.GetRoleByName)

	ar.POST("/users", adminHandler.CreateUser)
	ar.GET("/users", adminHandler.GetAllUsers)
}
