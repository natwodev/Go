package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/nguyenhuynhnam/go-backend-pro/internal/app/model"
	"github.com/nguyenhuynhnam/go-backend-pro/internal/app/service"
)

type InventoryController struct {
	service service.InventoryService
}

func NewInventoryController(service service.InventoryService) *InventoryController {
	return &InventoryController{service: service}
}

func (c *InventoryController) Create(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ProductID uint   `json:"product_id"`
		Quantity  int    `json:"quantity"`
		Location  string `json:"location"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		c.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	inventory, err := c.service.CreateInventory(input.ProductID, input.Quantity, input.Location)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "record not found") {
			c.respondWithError(w, http.StatusBadRequest, "product_id does not exist")
			return
		}
		c.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	c.respondWithJSON(w, http.StatusCreated, "Inventory created successfully", inventory)
}

func (c *InventoryController) GetAll(w http.ResponseWriter, r *http.Request) {
	inventories, err := c.service.GetAllInventories()
	if err != nil {
		c.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	c.respondWithJSON(w, http.StatusOK, "Inventories retrieved successfully", inventories)
}

func (c *InventoryController) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/inventories/")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.respondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	inventory, err := c.service.GetInventoryByID(uint(id))
	if err != nil {
		c.respondWithError(w, http.StatusNotFound, "Inventory not found")
		return
	}
	c.respondWithJSON(w, http.StatusOK, "Inventory retrieved successfully", inventory)
}

func (c *InventoryController) Increase(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/inventories/")
	idStr = strings.TrimSuffix(idStr, "/increase")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.respondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var input struct {
		Amount int `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		c.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	inventory, err := c.service.IncreaseInventory(uint(id), input.Amount)
	if err != nil {
		if strings.Contains(err.Error(), "amount must be greater than zero") {
			c.respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		c.respondWithError(w, http.StatusNotFound, "Inventory not found")
		return
	}
	c.respondWithJSON(w, http.StatusOK, "Inventory increased successfully", inventory)
}

func (c *InventoryController) Decrease(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/inventories/")
	idStr = strings.TrimSuffix(idStr, "/decrease")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.respondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var input struct {
		Amount int `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		c.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	inventory, err := c.service.DecreaseInventory(uint(id), input.Amount)
	if err != nil {
		if strings.Contains(err.Error(), "amount must be greater than zero") {
			c.respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		if errors.Is(err, service.ErrInventoryInsufficientQuantity) {
			c.respondWithError(w, http.StatusBadRequest, "insufficient quantity")
			return
		}
		c.respondWithError(w, http.StatusNotFound, "Inventory not found")
		return
	}
	c.respondWithJSON(w, http.StatusOK, "Inventory decreased successfully", inventory)
}

func (c *InventoryController) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/inventories/")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.respondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}
	if err := c.service.DeleteInventory(uint(id)); err != nil {
		c.respondWithError(w, http.StatusNotFound, "Inventory not found")
		return
	}
	c.respondWithJSON(w, http.StatusOK, "Inventory deleted successfully", nil)
}

func (c *InventoryController) respondWithError(w http.ResponseWriter, code int, message string) {
	c.respondWithJSON(w, code, message, nil)
}

func (c *InventoryController) respondWithJSON(w http.ResponseWriter, code int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(model.APIResponse{
		Status:  strconv.Itoa(code),
		Message: message,
		Data:    data,
	})
}
