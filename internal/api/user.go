package api

// tag binding
// binding 参考1：https://juejin.cn/post/6863765115456454664#heading-5
// binding 参考2：https://github.com/go-playground/validator

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
