package model

type GetUserByIdResponse struct {
	Name     string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Status   string `json:"status"`
}
