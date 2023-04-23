package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-auth/core/useCases"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqToken := c.GetHeader("authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]
		user, err := useCases.VerifyToken{}.Execute(reqToken)

		if err != nil {
			c.AbortWithStatusJSON(err.StatusCode, err.Error())
			return
		}
		c.Set("user", user)
		//auth.TokenValid(c)
		c.Next()
	}
}
