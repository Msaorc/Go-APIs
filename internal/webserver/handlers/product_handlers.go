package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Msaorc/Go-APIs/internal/dto"
	"github.com/Msaorc/Go-APIs/internal/entity"
	"github.com/Msaorc/Go-APIs/internal/infra/database"
	entityPKG "github.com/Msaorc/Go-APIs/pkg/entity"
	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

// Create product godoc
// @Summary      Create Product
// @Description  Create Product
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        request   body      dto.CreateProductInput  true  "product request"
// @Success      201
// @Failure      400  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /users/authenticate [post]
// @Security ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	err = h.ProductDB.Create(p)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// Get product godoc
// @Summary      Get Product
// @Description  Get Product
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        request   body      dto.CreateProductInput  true  "product request"
// @Success      200  {object}  entity.Product
// @Failure      400  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /users/authenticate [post]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}

// Create product godoc
// @Summary      Create Product
// @Description  Create Product
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        request   body      dto.CreateProductInput  true  "product request"
// @Success      201
// @Failure      400  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /users/authenticate [post]
// @Security ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product.ID, err = entityPKG.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = h.ProductDB.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Get All godoc
// @Summary      Get All Product
// @Description  Get All Product
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        page    query    string   false  "page number"
// @Param        limit   query    string   false   "limit"
// @Success      200     {array}  entity.Product
// @Failure      404  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /products [get]
// @Security ApiKeyAuth
func (h *ProductHandler) FindAllProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")

	pageint, err := strconv.Atoi(page)
	if err != nil {
		pageint = 0
	}
	limitint, err := strconv.Atoi(limit)
	if err != nil {
		limitint = 0
	}

	products, err := h.ProductDB.FindAll(pageint, limitint, sort)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		errorMessage := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// Create product godoc
// @Summary      Create Product
// @Description  Create Product
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        request   body      dto.CreateProductInput  true  "product request"
// @Success      201
// @Failure      400  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /users/authenticate [post]
// @Security ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = h.ProductDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
