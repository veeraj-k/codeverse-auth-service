package db

import (
	"auth-service/internal/auth"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	DB.Migrator().DropTable(&auth.User{}, &auth.Role{}, "user_roles")

	return DB
}

func MigrateDB() {

	DB.AutoMigrate(&auth.User{}, &auth.Role{})
}

func SeedDB() {
	userRepo := auth.NewUserRepository(DB)

	userRepo.CreateRole(&auth.Role{Name: "ADMIN"})
	userRepo.CreateRole(&auth.Role{Name: "SUPER_ADMIN"})
	userRepo.CreateRole(&auth.Role{Name: "USER"})

	hp, _ := bcrypt.GenerateFromPassword([]byte(os.Getenv("ADMIN_PASSWORD")), bcrypt.DefaultCost)

	role, _ := userRepo.GetRoleByName("ADMIN")

	userRepo.CreateUser(&auth.User{Username: os.Getenv("ADMIN_USERNAME"), Email: os.Getenv("ADMIN_EMAIL"), Password: string(hp), Role: []auth.Role{*role}})
}
