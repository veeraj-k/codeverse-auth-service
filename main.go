package main

import (
	"auth-service/config"
	"auth-service/db"
	"auth-service/internal/admin"
	"auth-service/internal/auth"
	"auth-service/internal/middlewares/pms"

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
	pms.RegisterRoute(r)

	r.Run("127.0.0.1:4000")

}
