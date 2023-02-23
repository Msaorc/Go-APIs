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
	configs, err := configs.LoadConfigs(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("file:APIgo.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.User{}, &entity.Product{})
	productHandler := handlers.NewProductHandler(database.NewProduct(db))
	userHandle := handlers.NewUserHandler(database.NewUser(db), configs.TokenAuth, configs.JwtExperesIn)
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Get("/products", productHandler.FindAllProducts)
	mux.Post("/products", productHandler.CreateProduct)
	mux.Get("/products/{id}", productHandler.GetProduct)
	mux.Put("/products/{id}", productHandler.UpdateProduct)
	mux.Delete("/products/{id}", productHandler.DeleteProduct)
	mux.Post("/users", userHandle.CreateUser)
	mux.Post("/users/authenticate", userHandle.Authentication)
	http.ListenAndServe(":8081", mux)
}
