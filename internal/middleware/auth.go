package middleware

import (
	"net/http"
	"strings"

	"github.com/0xDevvvvv/makerble/pkg/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Authorization Header Missing"})
			c.Abort()
			return
		}
		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		jwttoken := headerParts[1]
		claims, err := utils.ParseToken(jwttoken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		//this is to set the context with the user ID and role so that these process  maynot be repeated
		c.Set("username", claims.UserName)
		c.Set("role", claims.Role)

		c.Next()

	}
}
