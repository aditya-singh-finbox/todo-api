package middleware

import (
	"net/http"
	"strings"

	"github.com/aditya-singh-finbox/todo-api/internal/auth"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtService *auth.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "authorization header is required",
			})
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "nvalid authorization header format",
			})
			c.Abort()
			return
		}
		tokenstring := parts[1]
		if tokenstring == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token is required",
			})
			c.Abort()
			return
		}
		claims, err := jwtService.ValidateToken(tokenstring)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid or expired token",
			})
			c.Abort()
			return
		}
		c.Set("userID", claims.UserID)
		c.Next()
	}
}
