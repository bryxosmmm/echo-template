package models

type ClientSignUp struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required,min=3, max=100"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}

type ClientSignIn struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}

type SignSuccess struct {
	Token string `json:"token"`
	ID    string `json:"id"`
}
