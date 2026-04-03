package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/nguyenhuynhnam/go-backend-pro/internal/app/model"
	"github.com/nguyenhuynhnam/go-backend-pro/internal/app/service"
)

type ProductController struct {
	service service.ProductService
}

func NewProductController(service service.ProductService) *ProductController {
	return &ProductController{service: service}
}

func (c *ProductController) Create(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		CategoryID  uint    `json:"category_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		c.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	product, err := c.service.CreateProduct(input.Name, input.Description, input.Price, input.CategoryID)
	if err != nil {
		c.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	c.respondWithJSON(w, http.StatusCreated, "Product created successfully", product)
}

func (c *ProductController) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := c.service.GetAllProducts()
	if err != nil {
		c.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	c.respondWithJSON(w, http.StatusOK, "Products retrieved successfully", products)
}

func (c *ProductController) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/products/")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.respondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	product, err := c.service.GetProductByID(uint(id))
	if err != nil {
		c.respondWithError(w, http.StatusNotFound, "Product not found")
		return
	}
	c.respondWithJSON(w, http.StatusOK, "Product retrieved successfully", product)
}

func (c *ProductController) Update(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/products/")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.respondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var input struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		CategoryID  uint    `json:"category_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		c.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	product, err := c.service.UpdateProduct(uint(id), input.Name, input.Description, input.Price, input.CategoryID)
	if err != nil {
		c.respondWithError(w, http.StatusNotFound, "Product not found")
		return
	}
	c.respondWithJSON(w, http.StatusOK, "Product updated successfully", product)
}

func (c *ProductController) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/products/")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.respondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}
	if err := c.service.DeleteProduct(uint(id)); err != nil {
		c.respondWithError(w, http.StatusNotFound, "Product not found")
		return
	}
	c.respondWithJSON(w, http.StatusOK, "Product deleted successfully", nil)
}

func (c *ProductController) respondWithError(w http.ResponseWriter, code int, message string) {
	c.respondWithJSON(w, code, message, nil)
}

func (c *ProductController) respondWithJSON(w http.ResponseWriter, code int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(model.APIResponse{
		Status:  strconv.Itoa(code),
		Message: message,
		Data:    data,
	})
}
