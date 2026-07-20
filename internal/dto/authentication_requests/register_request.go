package dto

type RegisterRequest struct {
	UserName    string `json:"user_name" binding:"required"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Password    string `json:"password" binding:"required"`
}