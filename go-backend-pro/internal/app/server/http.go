package server

import (
	"net/http"
	"strings"
	"time"

	"github.com/nguyenhuynhnam/go-backend-pro/internal/app/controller"
	"github.com/nguyenhuynhnam/go-backend-pro/internal/app/handler"
)

func NewHTTPServer(
	addr string,
	categoryCtrl *controller.CategoryController,
	productCtrl *controller.ProductController,
	inventoryCtrl *controller.InventoryController,
) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handler.Health)
	mux.HandleFunc("/api/v1/ping", handler.Ping)

	mux.HandleFunc("/api/v1/categories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			categoryCtrl.GetAll(w, r)
		case http.MethodPost:
			categoryCtrl.Create(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/api/v1/categories/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			categoryCtrl.GetByID(w, r)
		case http.MethodPut:
			categoryCtrl.Update(w, r)
		case http.MethodDelete:
			categoryCtrl.Delete(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/api/v1/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			productCtrl.GetAll(w, r)
		case http.MethodPost:
			productCtrl.Create(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/api/v1/products/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			productCtrl.GetByID(w, r)
		case http.MethodPut:
			productCtrl.Update(w, r)
		case http.MethodDelete:
			productCtrl.Delete(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/api/v1/inventories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			inventoryCtrl.GetAll(w, r)
		case http.MethodPost:
			inventoryCtrl.Create(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/api/v1/inventories/", func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodPost && strings.HasSuffix(r.URL.Path, "/increase"):
			inventoryCtrl.Increase(w, r)
		case r.Method == http.MethodPost && strings.HasSuffix(r.URL.Path, "/decrease"):
			inventoryCtrl.Decrease(w, r)
		case r.Method == http.MethodGet:
			inventoryCtrl.GetByID(w, r)
		case r.Method == http.MethodDelete:
			inventoryCtrl.Delete(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
}
