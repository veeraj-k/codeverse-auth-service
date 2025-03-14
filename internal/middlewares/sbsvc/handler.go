package sbsvc

import (
	"auth-service/internal/admin"
	"auth-service/internal/utils"
	"os"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(r *gin.Engine) {

	pmr := r.Group("/api/submission")

	requiredRoles := map[string][]string{
		"GET":    {"SUPER_ADMIN", "ADMIN", "USER"},
		"POST":   {"SUPER_ADMIN", "ADMIN", "USER"},
		"PUT":    {"SUPER_ADMIN", "ADMIN"},
		"DELETE": {"SUPER_ADMIN", "ADMIN"},
	}

	pmr.Use(admin.AdminAuthMiddleware(requiredRoles))

	pmr.Any("/*path", func(c *gin.Context) {
		utils.ForwardRequestV2(c, os.Getenv("SUBMISSION_SERVICE_URL"))
	})

}
