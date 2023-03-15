package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Msaorc/Go-APIs/internal/dto"
	"github.com/Msaorc/Go-APIs/internal/entity"
	"github.com/Msaorc/Go-APIs/internal/infra/database"
	"github.com/go-chi/jwtauth"
)

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(userDB database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB: userDB,
	}
}

// Authenticate user godoc
// @Summary      Authenticate User
// @Description  Authenticate User
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request   body      dto.AuthenticationUserInput  true  "authenticate request"
// @Success      200  {object}  dto.AuthenticationUserOutput
// @Failure      400  {object}  dto.Error
// @Failure      401  {object}  dto.Error
// @Failure      404  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /users/authenticate [post]
func (uh *UserHandler) Authentication(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExperiesIn := r.Context().Value("experesIn").(int)
	var user dto.AuthenticationUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	u, err := uh.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		errorMessage := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	if !u.ValidatePassword(user.Password) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		errorMessage := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub":  u.ID.String(),
		"exp":  time.Now().Add(time.Second * time.Duration(jwtExperiesIn)).Unix(),
		"name": u.Name,
	})

	accessToken := dto.AuthenticationUserOutput{AccessToken: tokenString}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// Create user godoc
// @Summary      Create User
// @Description  Create User
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request   body      dto.CreateUserInput  true  "user request"
// @Success      201
// @Failure      400  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /users [post]
func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	err = uh.UserDB.Create(u)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
