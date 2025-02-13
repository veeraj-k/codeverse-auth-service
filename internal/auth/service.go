package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo UserRepository
}

func NewAuthService(repo UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(req RegisterRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}
	req.Password = string(hashedPassword)

	userRole, err := s.repo.GetRoleByName("USER")

	if err != nil {
		return err
	}

	user := User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Role:     []Role{*userRole},
	}
	return s.repo.CreateUser(&user)
}

func (s *AuthService) Login(req LoginRequest) (AuthResponse, error) {
	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return AuthResponse{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return AuthResponse{}, errors.New("invalid password")
	}

	roles := []string{}

	for _, role := range user.Role {
		roles = append(roles, role.Name)
	}
	token, err := GenerateToken(user, roles)
	if err != nil {
		return AuthResponse{}, err
	}
	return AuthResponse{UserID: user.ID, Token: token}, nil
}
