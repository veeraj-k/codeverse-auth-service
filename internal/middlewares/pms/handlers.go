package pms

import (
	"auth-service/internal/admin"
	"os"

	"auth-service/internal/utils"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(r *gin.Engine) {

	pmr := r.Group("/api/problems")

	requiredRoles := map[string][]string{
		"GET":    {"SUPER_ADMIN", "ADMIN", "USER"},
		"POST":   {"SUPER_ADMIN", "ADMIN"},
		"PUT":    {"SUPER_ADMIN", "ADMIN"},
		"DELETE": {"SUPER_ADMIN", "ADMIN"},
	}

	pmr.Use(admin.AdminAuthMiddleware(requiredRoles))

	pmr.Any("/*path", func(c *gin.Context) {
		utils.ForwardRequest(c, os.Getenv("PMS_SERVICE_URL"))
	})

}
