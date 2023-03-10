package main

import (
	"net/http"

	"github.com/Msaorc/Go-APIs/configs"
	"github.com/Msaorc/Go-APIs/internal/entity"
	"github.com/Msaorc/Go-APIs/internal/infra/database"
	"github.com/Msaorc/Go-APIs/internal/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
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
	userHandle := handlers.NewUserHandler(database.NewUser(db))
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.WithValue("jwt", configs.TokenAuth))
	mux.Use(middleware.WithValue("experesIn", configs.JwtExperesIn))

	mux.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Get("/", productHandler.FindAllProducts)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	mux.Post("/users", userHandle.CreateUser)
	mux.Post("/users/authenticate", userHandle.Authentication)
	http.ListenAndServe(":8081", mux)
}
