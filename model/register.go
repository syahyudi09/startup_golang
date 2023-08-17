package model

type RegisterUserInput struct{
	Name string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Role string 
}

type LoginUser struct{
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CheckEmailAvailable struct {
	Email string `json:"email" binding:"required,email"`
}