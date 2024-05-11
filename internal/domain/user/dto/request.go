package dto

import "github.com/go-playground/validator/v10"

type UserCreateRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (r *UserCreateRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
