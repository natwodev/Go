package main

import (
	"log"
	"os"

	"github.com/nguyenhuynhnam/go-backend-pro/internal/app/controller"
	"github.com/nguyenhuynhnam/go-backend-pro/internal/app/model"
	"github.com/nguyenhuynhnam/go-backend-pro/internal/app/repository"
	"github.com/nguyenhuynhnam/go-backend-pro/internal/app/server"
	"github.com/nguyenhuynhnam/go-backend-pro/internal/app/service"
	"github.com/nguyenhuynhnam/go-backend-pro/internal/pkg/postgres"
)

func main() {
	// 1. Initialize Database
	db, err := postgres.NewConnection()
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	// 1b. Auto Migration
	log.Println("running auto migration...")
	db.AutoMigrate(&model.Category{}, &model.Product{}, &model.Inventory{})

	// 2. Initialize Repository, Service, Controller
	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryCtrl := controller.NewCategoryController(categoryService)
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productCtrl := controller.NewProductController(productService)
	inventoryRepo := repository.NewInventoryRepository(db)
	inventoryService := service.NewInventoryService(inventoryRepo, productRepo)
	inventoryCtrl := controller.NewInventoryController(inventoryService)

	// 3. Start HTTP Server
	addr := getenv("APP_ADDR", ":8080")
	httpServer := server.NewHTTPServer(addr, categoryCtrl, productCtrl, inventoryCtrl)

	log.Printf("server listening on %s", addr)
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

func getenv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
