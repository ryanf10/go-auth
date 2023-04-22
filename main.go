package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/joho/godotenv"
	"go-auth/infrastructure/controllers"
	"go-auth/infrastructure/controllers/forms"
	"net/http"
)

func main() {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		fmt.Println(errEnv)
		return
	}
	r := gin.Default()
	binding.Validator = new(forms.DefaultValidator)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	userController := controllers.NewUserController()
	r.POST("/register", userController.Register)

	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
