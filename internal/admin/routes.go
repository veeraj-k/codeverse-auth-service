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
	requiredRoles := map[string][]string{
		"GET":    {"SUPER_ADMIN", "ADMIN"},
		"POST":   {"SUPER_ADMIN", "ADMIN"},
		"PUT":    {"SUPER_ADMIN", "ADMIN"},
		"DELETE": {"SUPER_ADMIN", "ADMIN"},
	}

	ar := r.Group("/api/admin", AdminAuthMiddleware(requiredRoles))

	ar.POST("/roles", adminHandler.CreateRole)
	ar.GET("/roles", adminHandler.GetRoles)
	ar.DELETE("/roles/:name", adminHandler.DeleteRoleByName)
	ar.GET("/roles/:name", adminHandler.GetRoleByName)

	ar.POST("/users", adminHandler.CreateUser)
	ar.GET("/users", adminHandler.GetAllUsers)
}
