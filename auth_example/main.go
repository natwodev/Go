package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/natwodev/Go/auth"
)

// AuthMiddleware kiểm tra JWT token trong header
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}

		claims, err := auth.VerifyJWT(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}

func main() {
	r := gin.Default()

	// Route để lấy token
	r.POST("/login", func(c *gin.Context) {
		var loginVals struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&loginVals); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		// Giả sử kiểm tra login thành công
		token, err := auth.GenerateJWT(loginVals.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	// Group các route cần xác thực
	authorized := r.Group("/api")
	authorized.Use(AuthMiddleware())
	{
		authorized.GET("/profile", func(c *gin.Context) {
			username, _ := c.Get("username")
			c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Hello %s, this is your profile", username)})
		})
	}

	r.Run(":8080")
}
