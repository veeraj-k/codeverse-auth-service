package admin

type CreateUserRequest struct {
	Username string   `json:"username" binding:"required"`
	Email    string   `json:"email" binding:"required"`
	Password string   `json:"password" binding:"required"`
	Role     []string `json:"role"`
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
