package auth

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateRoleRequest struct {
	Name string `json:"name" binding:"required"`
}

type UserResponse struct {
	ID       uint           `json:"id"`
	Username string         `json:"username"`
	Email    string         `json:"email"`
	Role     []RoleResponse `json:"role"`
}
type RoleResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
type AuthResponse struct {
	UserID uint   `json:"user_id"`
	Token  string `json:"token"`
}
