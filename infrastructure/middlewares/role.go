package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-auth/core/entities"
	"net/http"
)

func RoleMiddleware(rolesName []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		value, exists := c.Get("user")
		user := value.(entities.User)

		if exists == false {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		} else {
			found := false
			for _, name := range rolesName {
				if name == user.Role.Name {
					found = true
					break
				}
			}
			if !found {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
				return
			}
		}
		c.Next()
	}
}
