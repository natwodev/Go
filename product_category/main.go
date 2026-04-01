package main

import (
	"net/http"
	"strconv"

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

// In-memory "Database"
var (
	categories = []Category{
		{ID: 1, Name: "Electronics"},
		{ID: 2, Name: "Books"},
	}
	nextCategoryID = 3
)

func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		categoriesGroup := v1.Group("/categories")
		{
			categoriesGroup.GET("", getCategories)
			categoriesGroup.GET("/:id", getCategory)
			categoriesGroup.POST("", createCategory)
			categoriesGroup.PUT("/:id", updateCategory)
			categoriesGroup.DELETE("/:id", deleteCategory)
		}
	}

	r.Run(":8080")
}

// --- Category Handlers ---

func getCategories(c *gin.Context) {
	c.JSON(http.StatusOK, categories)
}

func getCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, cat := range categories {
		if cat.ID == id {
			c.JSON(http.StatusOK, cat)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
}

func createCategory(c *gin.Context) {
	var newCategory Category
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newCategory.ID = nextCategoryID
	nextCategoryID++
	categories = append(categories, newCategory)
	c.JSON(http.StatusCreated, newCategory)
}

func updateCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedCategory Category
	if err := c.ShouldBindJSON(&updatedCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, cat := range categories {
		if cat.ID == id {
			updatedCategory.ID = id
			categories[i] = updatedCategory
			c.JSON(http.StatusOK, updatedCategory)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
}

func deleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, cat := range categories {
		if cat.ID == id {
			categories = append(categories[:i], categories[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
}
