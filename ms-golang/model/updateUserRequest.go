package model

type UpdateUserRequest struct {
	Id       string `uri:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Status   string `json:"status" binding:"oneof=ENABLED DISABLED" message:"400.002=oneof"`
}
