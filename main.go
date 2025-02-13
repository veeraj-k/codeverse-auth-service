package main

import (
	"auth-service/config"
	"auth-service/db"
	"auth-service/internal/admin"
	"auth-service/internal/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	config.Loadenv()
	DB = db.InitDB()
	db.MigrateDB()
	db.SeedDB()

}

func main() {

	r := gin.Default()

	auth.RegisterRoutes(r, DB)
	admin.RegisterRoutes(r, DB)

	r.Run("127.0.0.1:8080")

}
