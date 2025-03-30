package main

import (
	"auth-service/config"
	"auth-service/db"
	"auth-service/internal/admin"
	"auth-service/internal/auth"
	"auth-service/internal/middlewares/pms"
	"auth-service/internal/middlewares/sbsvc"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	if os.Getenv("LOAD_ENV") == "true" {
		config.Loadenv()
	}
	DB = db.InitDB()
	db.MigrateDB()
	db.SeedDB()

}

func main() {

	r := gin.Default()

	r.Handle("GET", "/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	auth.RegisterRoutes(r, DB)
	admin.RegisterRoutes(r, DB)
	pms.RegisterRoute(r)
	sbsvc.RegisterRoute(r)

	r.Run("0.0.0.0:4000")

}
