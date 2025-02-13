package main

import (
	"auth-service/config"
	"auth-service/internal/admin"
	"auth-service/internal/auth"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	config.Loadenv()
}

var DB *gorm.DB

func main() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB.Migrator().DropTable(&auth.User{}, &auth.Role{}, "user_roles")

	DB.AutoMigrate(&auth.User{}, &auth.Role{})
	r := gin.Default()

	userRepo := auth.NewUserRepository(DB)

	userRepo.CreateRole(&auth.Role{Name: "ADMIN"})
	userRepo.CreateRole(&auth.Role{Name: "SUPER_ADMIN"})
	userRepo.CreateRole(&auth.Role{Name: "USER"})

	hp, _ := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)

	role, _ := userRepo.GetRoleByName("ADMIN")

	userRepo.CreateUser(&auth.User{Username: "admin", Email: "admin", Password: string(hp), Role: []auth.Role{*role}})

	auth.RegisterRoutes(r, DB)
	admin.RegisterRoutes(r, DB)

	r.Run("127.0.0.1:8080")

}
