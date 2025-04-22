package middleware

import "github.com/gin-gonic/gin"

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func RoleMiddleware(role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
