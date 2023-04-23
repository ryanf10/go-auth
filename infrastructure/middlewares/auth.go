package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-auth/core/useCases"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqToken := c.GetHeader("authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		if len(splitToken) < 2 {
			c.AbortWithStatusJSON(http.StatusForbidden, "Unauthorized")
			return
		}
		reqToken = splitToken[1]
		user, err := useCases.VerifyToken{}.Execute(reqToken)

		if err != nil {
			c.AbortWithStatusJSON(err.StatusCode, err.Error())
			return
		}
		c.Set("user", user)

		c.Next()
	}
}
