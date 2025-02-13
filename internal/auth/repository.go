package auth

import "gorm.io/gorm"

type UserRepository interface {
	CreateUser(user *User) error
	GetUserByEmail(email string) (*User, error)
	GetAllUsers() ([]User, error)
	DeleteUser(id uint) error
	UpdateUser(user *User) error

	CreateRole(role *Role) error
	GetRoleByName(name string) (*Role, error)
	GetAllRoles() ([]Role, error)
	DeleteRoleByName(name string) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	err := r.DB.Preload("Role").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) CreateUser(user *User) error {
	return r.DB.Create(user).Error
}

func (r *userRepository) GetUserByEmail(email string) (*User, error) {
	var user User
	err := r.DB.Preload("Role").Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) DeleteUser(id uint) error {
	return r.DB.Delete(&User{}, id).Error
}

func (r *userRepository) UpdateUser(user *User) error {
	return r.DB.Save(user).Error
}

func (r *userRepository) CreateRole(role *Role) error {
	return r.DB.Create(role).Error
}

func (r *userRepository) GetRoleByName(name string) (*Role, error) {
	var role Role
	err := r.DB.Where("name = ?", name).First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *userRepository) GetAllRoles() ([]Role, error) {
	var roles []Role
	err := r.DB.Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *userRepository) DeleteRoleByName(name string) error {
	role, err := r.GetRoleByName(name)
	if err != nil {
		return err
	}
	return r.DB.Delete(role).Error
}
