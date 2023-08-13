package model

type RegisterUserInput struct{
	Name string 
	Occupation string
	Email string
	Password string
	Role string
	Token string
}

type LoginUser struct{
	Email string `json:"email"`
	Password string `json:"password"`
}

type CheckEmailAvalible struct {
	Email string `json:"email"`
}