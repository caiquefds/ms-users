package model

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required" message:"400.001=required"`
	Email    string `json:"email" binding:"required" message:"400.001=required"`
	Password string `json:"password" binding:"required" message:"400.001=required"`
}
