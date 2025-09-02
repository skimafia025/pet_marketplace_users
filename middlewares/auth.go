package middlewares

import (
	"net/http"
	"pet_marketplace_users/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")

		if authHeader == "" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization token not provided",
			})
			context.Abort()
			return
		}
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token format",
			})
			context.Abort()
			return
		}

		token := tokenParts[1]

		claims, err := utils.ValidateToken(token)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid Token",
			})
			context.Abort()
			return
		}

		context.Set("user_id", claims.UserID)
		context.Set("user_email", claims.Email)

		context.Next()
	}
}
