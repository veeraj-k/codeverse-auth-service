package admin

import (
	"auth-service/internal/auth"

	"golang.org/x/crypto/bcrypt"
)

type AdminAuthService struct {
	repo auth.UserRepository
}

func NewAdminAuthService(repo auth.UserRepository) *AdminAuthService {
	return &AdminAuthService{repo: repo}
}

func (s *AdminAuthService) CreateRole(req CreateRoleRequest) error {
	role := auth.Role{
		Name: req.Name,
	}
	return s.repo.CreateRole(&role)
}

func (s *AdminAuthService) GetRoles() ([]RoleResponse, error) {
	var rolesResponse []RoleResponse
	roles, err := s.repo.GetAllRoles()
	if err != nil {
		return nil, err
	}

	for _, role := range roles {
		rolesResponse = append(rolesResponse, RoleResponse{
			ID:   role.ID,
			Name: role.Name})
	}

	return rolesResponse, nil
}

func (s *AdminAuthService) GetRoleByName(name string) (*RoleResponse, error) {
	role, err := s.repo.GetRoleByName(name)
	if err != nil {
		return nil, err
	}
	return &RoleResponse{
		ID:   role.ID,
		Name: role.Name,
	}, nil
}

func (s *AdminAuthService) DeleteRoleByName(name string) error {

	return s.repo.DeleteRoleByName(name)
}

func (s *AdminAuthService) CreateUser(req CreateUserRequest) error {

	roles := []auth.Role{}
	for _, roleName := range req.Role {
		role, err := s.repo.GetRoleByName(roleName)
		if err != nil {
			return err
		}
		roles = append(roles, *role)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	req.Password = string(hashedPassword)
	user := auth.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Role:     roles,
	}
	return s.repo.CreateUser(&user)
}

func (s *AdminAuthService) GetUsers() ([]UserResponse, error) {
	var usersResponse []UserResponse
	users, err := s.repo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		roles := []RoleResponse{}
		for _, role := range user.Role {
			roles = append(roles, RoleResponse{
				ID:   role.ID,
				Name: role.Name,
			})
		}
		usersResponse = append(usersResponse, UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Role:     roles,
		})
	}

	return usersResponse, nil
}
