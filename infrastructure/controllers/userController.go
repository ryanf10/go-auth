package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-auth/core/useCases"
	"go-auth/infrastructure/controllers/forms"
	"go-auth/infrastructure/repostiories"
	"net/http"
)

type UserController struct {
	createUser *useCases.CreateUser
}

var registerFormValidator = new(forms.RegisterFormValidator)

func NewUserController() *UserController {
	p := new(UserController)
	userRepository := repostiories.NewUserRepository()
	p.createUser = useCases.NewCreateUser(userRepository)
	return p
}

func (userController UserController) Register(c *gin.Context) {
	var registerForm forms.RegisterForm
	if validationErr := c.ShouldBindJSON(&registerForm); validationErr != nil {
		fmt.Println(validationErr)
		message := registerFormValidator.ErrorHandler(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}
	user := userController.createUser.Execute(registerForm.Name, registerForm.Email, registerForm.Password)
	c.JSON(http.StatusOK, user)

}
