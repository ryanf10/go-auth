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
	login      *useCases.Login
}

var registerFormValidator = new(forms.RegisterFormValidator)
var loginFormValidator = new(forms.LoginFormValidator)

func NewUserController() *UserController {
	p := new(UserController)
	userRepository := repostiories.NewUserRepository()
	roleRepository := repostiories.NewRoleRepository()
	p.createUser = useCases.NewCreateUser(userRepository, roleRepository)
	p.login = useCases.NewLogin(userRepository)
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

func (userController UserController) Login(c *gin.Context) {

	var loginForm forms.LoginForm
	if validationErr := c.ShouldBindJSON(&loginForm); validationErr != nil {
		fmt.Println(validationErr)
		message := loginFormValidator.ErrorHandler(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}
	user, token, err := userController.login.Execute(loginForm.Email, loginForm.Password)
	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode, gin.H{"status": "error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"token": token, "user": user}})

}
