package api

type RegisterRequest struct {
	UserName string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type DeleteRequest struct {
	Id int `json:"id" binding:"required"`
}

type UpdateRequest struct {
	Id       int    `json:"id" binding:"required"`
	UserName string `json:"username" binding:"required"`
}

type InfoRequest struct {
	Id int `json:"id" binding:"required"`
}
