package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/joho/godotenv"
	"go-auth/infrastructure/controllers"
	"go-auth/infrastructure/controllers/forms"
	"go-auth/infrastructure/middlewares"
	"net/http"
)

func main() {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		fmt.Println(errEnv)
	}
	r := gin.Default()
	binding.Validator = new(forms.DefaultValidator)

	r.GET("/ping", controllers.PingController{}.Ping)

	v1 := r.Group("/api/v1")
	{
		userController := controllers.NewUserController()
		v1.POST("/register", userController.Register)
		v1.POST("/login", userController.Login)

		// auth route
		auth := v1.Group("/")
		auth.Use(middlewares.AuthMiddleware())
		{
			auth.GET("/private", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"status": "success",
					"data":   gin.H{},
				})
			})

			role := auth.Group("/")
			{
				roleAdmin := role.Group("/")
				roleAdmin.Use(middlewares.RoleMiddleware([]string{"admin"}))
				{
					roleAdmin.GET("/private/admin", func(c *gin.Context) {
						c.JSON(http.StatusOK, gin.H{
							"status": "success",
							"data":   gin.H{},
						})
					})
				}
			}
		}
	}

	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
