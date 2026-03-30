package main

import (
	"github.com/gin-gonic/gin"
)

// Models
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name" binding:"required"`
}

type Product struct {
	ID         int     `json:"id"`
	Name       string  `json:"name" binding:"required"`
	Price      float64 `json:"price" binding:"required"`
	CategoryID int     `json:"category_id" binding:"required"`
}

func main() {
	r := gin.Default()
	r.Run(":8080")
}
