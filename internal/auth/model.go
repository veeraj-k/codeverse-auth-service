package auth

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Email    string `gorm:"unique"`
	Password string
	Role     []Role `gorm:"many2many:user_roles;"`
}

type Role struct {
	gorm.Model
	Name string `gorm:"unique"`
}
