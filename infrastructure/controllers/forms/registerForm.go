package forms

import "github.com/go-playground/validator/v10"

type RegisterForm struct {
	Name            string `form:"name" json:"name" binding:"required,min=3,max=20,name"`
	Email           string `form:"email" json:"email" binding:"required,email"`
	Password        string `form:"password" json:"password" binding:"required,min=3,max=50,eqfield=ConfirmPassword"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" binding:"required,min=3,max=50"`
}

type RegisterFormValidator struct{}

func (f RegisterFormValidator) Name(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Field name is required"
		}
		return errMsg[0]
	case "min", "max":
		return "Field name should be between 3 to 20 characters"
	case "name":
		return "Field name should not include any special characters or numbers"
	default:
		return "Something went wrong, please try again later"
	}
}

// Email ...
func (f RegisterFormValidator) Email(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Field email is required"
		}
		return errMsg[0]
	case "min", "max", "email":
		return "Field email is invalid"
	default:
		return "Something went wrong, please try again later"
	}
}

// Password ...
func (f RegisterFormValidator) Password(tag string) (message string) {
	switch tag {
	case "required":
		return "Field password is required"
	case "min", "max":
		return "Field password should be between 3 and 50 characters"
	case "eqfield":
		return "Field passwords does not match"
	default:
		return "Something went wrong, please try again later"
	}
}

func (f RegisterFormValidator) ErrorHandler(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:

		//if _, ok := err.(*json.UnmarshalTypeError); ok {
		//	return "Something went wrong, please try again later"
		//}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Name" {
				return f.Name(err.Tag())
			}

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
