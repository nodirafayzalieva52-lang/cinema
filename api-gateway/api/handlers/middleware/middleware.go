package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nodirafayzalieva52-lang/cinema/api-gateway/models/permission.go"
)

func RBAC(requiredPermission string) gin.HandlerFunc {

	return func(c *gin.Context) {
		role := c.GetString("role")
		if role == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		permissions, ok := permission.RolePermission[role]
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		if _, ok = permissions[requiredPermission]; !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Permission denied",
			})
			return
		}
		c.Next()
	}
}
