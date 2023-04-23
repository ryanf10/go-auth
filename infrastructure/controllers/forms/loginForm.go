package forms

import "github.com/go-playground/validator/v10"

type LoginForm struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}

type LoginFormValidator struct{}

// Email ...
func (f LoginFormValidator) Email(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Field email is required"
		}
		return errMsg[0]
	case "email":
		return "Field email is invalid"
	default:
		return "Something went wrong, please try again later"
	}
}

// Password ...
func (f LoginFormValidator) Password(tag string) (message string) {
	switch tag {
	case "required":
		return "Field password is required"
	default:
		return "Something went wrong, please try again later"
	}
}

func (f LoginFormValidator) ErrorHandler(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:

		//if _, ok := err.(*json.UnmarshalTypeError); ok {
		//	return "Something went wrong, please try again later"
		//}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Email" {
				return f.Email(err.Tag())
			}

			if err.Field() == "Password" {
				return f.Password(err.Tag())
			}

		}
	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}
