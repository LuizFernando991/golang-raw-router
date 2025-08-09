package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/LuizFernando991/golang-api/infra/config"
)

type UserController struct {
	logger config.Logger
}

func NewUserController() *UserController {
	return &UserController{
		logger: *config.GetLogger("user-controller"),
	}
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	uc.logger.Info("create_user_route")

	type CreateUserRequest struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	var reqBody CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Mock de um usuário criado
	user := map[string]interface{}{
		"id":    1,
		"name":  reqBody.Name,
		"email": reqBody.Email,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
