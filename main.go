package main

import (
	"auth-service/config"
	"auth-service/db"
	"auth-service/internal/admin"
	"auth-service/internal/auth"
	"auth-service/internal/middlewares/pms"
	"auth-service/internal/middlewares/sbsvc"
	"os"
	"time"

	"github.com/gin-contrib/cors"
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
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
