package main

import (
	"net/http"

	"github.com/Msaorc/Go-APIs/configs"
	"github.com/Msaorc/Go-APIs/internal/entity"
	"github.com/Msaorc/Go-APIs/internal/handlers"
	"github.com/Msaorc/Go-APIs/internal/infra/database"
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
	mux := http.NewServeMux()
	mux.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8081", mux)
}
