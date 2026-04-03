package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/nguyenhuynhnam/go-backend-pro/internal/app/model"
	"github.com/nguyenhuynhnam/go-backend-pro/internal/app/service"
)

type CategoryController struct {
	service service.CategoryService
}

func NewCategoryController(service service.CategoryService) *CategoryController {
	return &CategoryController{service: service}
}

func (c *CategoryController) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		c.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	category, err := c.service.CreateCategory(input.Name)
	if err != nil {
		c.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	c.respondWithJSON(w, http.StatusCreated, "Category created successfully", category)
}

func (c *CategoryController) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	categories, err := c.service.GetAllCategories()
	if err != nil {
		c.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	c.respondWithJSON(w, http.StatusOK, "Categories retrieved successfully", categories)
}

func (c *CategoryController) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/categories/")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.respondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	category, err := c.service.GetCategoryByID(uint(id))
	if err != nil {
		c.respondWithError(w, http.StatusNotFound, "Category not found")
		return
	}

	c.respondWithJSON(w, http.StatusOK, "Category retrieved successfully", category)
}

func (c *CategoryController) Update(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/categories/")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.respondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var input struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		c.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	category, err := c.service.UpdateCategory(uint(id), input.Name)
	if err != nil {
		c.respondWithError(w, http.StatusNotFound, "Category not found")
		return
	}

	c.respondWithJSON(w, http.StatusOK, "Category updated successfully", category)
}

func (c *CategoryController) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/categories/")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.respondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := c.service.DeleteCategory(uint(id)); err != nil {
		c.respondWithError(w, http.StatusNotFound, "Category not found")
		return
	}

	c.respondWithJSON(w, http.StatusOK, "Category deleted successfully", nil)
}

func (c *CategoryController) respondWithError(w http.ResponseWriter, code int, message string) {
	c.respondWithJSON(w, code, message, nil)
}

func (c *CategoryController) respondWithJSON(w http.ResponseWriter, code int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(model.APIResponse{
		Status:  strconv.Itoa(code),
		Message: message,
		Data:    data,
	})
}
