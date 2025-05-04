package validator

import (
	"go-rest-api/model"

	"github.com/go-ozzo/ozzo-validation/v4/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)


type IUserValidator interface {
	UserValidate(user model.User) error
}

type userValidator struct{}		

func NewUserValidator() IUserValidator {
	return &userValidator{}
}	

func (v *userValidator) UserValidate(user model.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("email is required"),
			validation.RuneLength(1, 30).Error("email must be less than 30 characters"),
			is.Email.Error("email is not valid"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(6, 30).Error("password must be less than 30 characters"),
		),
		
	)
}

