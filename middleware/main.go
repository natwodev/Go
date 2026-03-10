package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// CustomLogger là một middleware ghi lại log yêu cầu
func CustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Xử lý yêu cầu
		c.Next()

		// Sau khi yêu cầu được xử lý
		latency := time.Since(t)
		status := c.Writer.Status()
		fmt.Printf("[%s] %d | %s | %s | %s\n",
			t.Format("2006/01/02 - 15:04:05"),
			status,
			latency,
			c.Request.Method,
			c.Request.URL.Path,
		)
	}
}

func main() {
	// Ghi log vào file
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.New()

	// Sử dụng Middleware Recovery chuẩn của Gin
	r.Use(gin.Recovery())
	
	// Sử dụng Logger tùy chỉnh
	r.Use(CustomLogger())

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.GET("/panic", func(c *gin.Context) {
		// Test recovery middleware
		panic("Oops, something went wrong!")
	})

	r.Run(":8080")
}
