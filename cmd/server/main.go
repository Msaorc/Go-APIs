package main

import (
	"net/http"

	"github.com/Msaorc/Go-APIs/configs"
	"github.com/Msaorc/Go-APIs/internal/entity"
	"github.com/Msaorc/Go-APIs/internal/infra/database"
	"github.com/Msaorc/Go-APIs/internal/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfigs(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("file:APIgo.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.User{}, &entity.Product{})
	productHandler := handlers.NewProductHandler(database.NewProduct(db))
	routers := chi.NewRouter()
	routers.Use(middleware.Logger)
	routers.Post("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8081", routers)
}
