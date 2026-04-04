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

type Inventory struct {
	ID        int    `json:"id"`
	ProductID int    `json:"product_id" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required"`
	Location  string `json:"location" binding:"required"`
}

type Order struct {
	ID           int     `json:"id"`
	ProductID    int     `json:"product_id" binding:"required"`
	Quantity     int     `json:"quantity" binding:"required"`
	TotalPrice   float64 `json:"total_price"`
	CustomerName string  `json:"customer_name" binding:"required"`
	Status       string  `json:"status"`
}

// In-memory "Database"
var (
	categories = []Category{
		{ID: 1, Name: "Electronics"},
		{ID: 2, Name: "Books"},
	}
	products = []Product{
		{ID: 1, Name: "Laptop", Price: 1200.0, CategoryID: 1},
		{ID: 2, Name: "Go Programming", Price: 45.0, CategoryID: 2},
	}
	inventories = []Inventory{
		{ID: 1, ProductID: 1, Quantity: 50, Location: "Warehouse A"},
		{ID: 2, ProductID: 2, Quantity: 200, Location: "Warehouse B"},
	}
	orders = []Order{
		{ID: 1, ProductID: 1, Quantity: 2, TotalPrice: 2400.0, CustomerName: "Nguyen Van A", Status: "pending"},
	}
	nextCategoryID  = 3
	nextProductID   = 3
	nextInventoryID = 3
	nextOrderID     = 2
)

func main() {
	r := gin.Default()

	// API Routes
	v1 := r.Group("/api/v1")
	{
		// Category Routes
		categoriesGroup := v1.Group("/categories")
		{
			categoriesGroup.GET("", getCategories)
			categoriesGroup.GET("/:id", getCategory)
			categoriesGroup.POST("", createCategory)
			categoriesGroup.PUT("/:id", updateCategory)
			categoriesGroup.DELETE("/:id", deleteCategory)
		}

		// Product Routes
		productsGroup := v1.Group("/products")
		{
			productsGroup.GET("", getProducts)
			productsGroup.GET("/:id", getProduct)
			productsGroup.POST("", createProduct)
			productsGroup.PUT("/:id", updateProduct)
			productsGroup.DELETE("/:id", deleteProduct)
		}

		// Inventory Routes
		inventoriesGroup := v1.Group("/inventories")
		{
			inventoriesGroup.GET("", getInventories)
			inventoriesGroup.GET("/:id", getInventory)
			inventoriesGroup.POST("", createInventory)
			inventoriesGroup.PUT("/:id", updateInventory)
			inventoriesGroup.DELETE("/:id", deleteInventory)
		}

		// Order Routes
		ordersGroup := v1.Group("/orders")
		{
			ordersGroup.GET("", getOrders)
			ordersGroup.GET("/:id", getOrder)
			ordersGroup.POST("", createOrder)
			ordersGroup.PUT("/:id", updateOrder)
			ordersGroup.DELETE("/:id", deleteOrder)
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

// --- Product Handlers ---

func getProducts(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

func getProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, p := range products {
		if p.ID == id {
			c.JSON(http.StatusOK, p)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}

func createProduct(c *gin.Context) {
	var newProduct Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simple check if category exists
	categoryExists := false
	for _, cat := range categories {
		if cat.ID == newProduct.CategoryID {
			categoryExists = true
			break
		}
	}
	if !categoryExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid CategoryID"})
		return
	}

	newProduct.ID = nextProductID
	nextProductID++
	products = append(products, newProduct)
	c.JSON(http.StatusCreated, newProduct)
}

func updateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedProduct Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, p := range products {
		if p.ID == id {
			updatedProduct.ID = id
			products[i] = updatedProduct
			c.JSON(http.StatusOK, updatedProduct)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}

func deleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}

// --- Inventory Handlers ---

func getInventories(c *gin.Context) {
	c.JSON(http.StatusOK, inventories)
}

func getInventory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, inv := range inventories {
		if inv.ID == id {
			c.JSON(http.StatusOK, inv)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
}

func createInventory(c *gin.Context) {
	var newInventory Inventory
	if err := c.ShouldBindJSON(&newInventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Kiểm tra product có tồn tại không
	productExists := false
	for _, p := range products {
		if p.ID == newInventory.ProductID {
			productExists = true
			break
		}
	}
	if !productExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product không tồn tại"})
		return
	}

	newInventory.ID = nextInventoryID
	nextInventoryID++
	inventories = append(inventories, newInventory)
	c.JSON(http.StatusCreated, newInventory)
}

func updateInventory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedInventory Inventory
	if err := c.ShouldBindJSON(&updatedInventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, inv := range inventories {
		if inv.ID == id {
			updatedInventory.ID = id
			inventories[i] = updatedInventory
			c.JSON(http.StatusOK, updatedInventory)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
}

func deleteInventory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, inv := range inventories {
		if inv.ID == id {
			inventories = append(inventories[:i], inventories[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Inventory deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
}

// --- Order Handlers ---

func getOrders(c *gin.Context) {
	c.JSON(http.StatusOK, orders)
}

func getOrder(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, o := range orders {
		if o.ID == id {
			c.JSON(http.StatusOK, o)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
}

func createOrder(c *gin.Context) {
	var newOrder Order
	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Kiểm tra product và tính tổng giá
	var productPrice float64
	productFound := false
	for _, p := range products {
		if p.ID == newOrder.ProductID {
			productPrice = p.Price
			productFound = true
			break
		}
	}
	if !productFound {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product không tồn tại"})
		return
	}

	newOrder.ID = nextOrderID
	nextOrderID++
	newOrder.TotalPrice = productPrice * float64(newOrder.Quantity)
	newOrder.Status = "pending"
	orders = append(orders, newOrder)
	c.JSON(http.StatusCreated, newOrder)
}

func updateOrder(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedOrder Order
	if err := c.ShouldBindJSON(&updatedOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, o := range orders {
		if o.ID == id {
			updatedOrder.ID = id
			orders[i] = updatedOrder
			c.JSON(http.StatusOK, updatedOrder)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
}

func deleteOrder(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, o := range orders {
		if o.ID == id {
			orders = append(orders[:i], orders[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Order deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
}
