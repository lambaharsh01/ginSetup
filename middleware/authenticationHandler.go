package middleware

import (
	"github.com/gin-gonic/gin"
	"ticketingBackend/utils"
	"strings"
	"net/http"
)

func AuthenticationHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success":false,
				"error": "Authorization header missing",
			})
            return
        }

        // Extract token from "Bearer <token>" format
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        claims, err := utils.ValidateJWT(tokenString)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success":false,
				"error": "Invalid or expired token",
			})
            return
        }

        // Attach claims to context
        c.Set("UserEmail", claims.UserEmail)
        c.Next()
    }
}