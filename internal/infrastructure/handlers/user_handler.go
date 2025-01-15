package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"to-do-list/internal/domain/models"
	"to-do-list/internal/domain/service"
	"to-do-list/internal/infrastructure/middleware"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		middleware.Response(w, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	if err := middleware.Validate.Struct(user); err != nil {
		middleware.Response(w, http.StatusBadRequest, fmt.Sprintf("Validation error: %v", err), nil)
		return
	}

	createUser, err := h.service.CreateUser(&user)
	if err != nil {
		middleware.Response(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	middleware.Response(w, http.StatusOK, "Successfully signup", map[string]interface{}{
		"id": createUser.UUID,
	})
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var user models.Login
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		middleware.Response(w, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	if err := middleware.Validate.Struct(user); err != nil {
		middleware.Response(w, http.StatusBadRequest, fmt.Sprintf("Validation error: %v", err), nil)
		return
	}

	if err := middleware.Validate.Struct(user); err != nil {
		middleware.Response(w, http.StatusBadRequest, fmt.Sprintf("Validation error: %v", err), nil)
		return
	}

	login, err := h.service.Login(&user)
	if err != nil {
		middleware.Response(w, http.StatusUnauthorized, err.Error(), nil)
		return
	}
	sess := middleware.CreateSession(w, r, login.UUID)

	err = middleware.CreateSession(w, r, login.UUID)
	fmt.Println("err::", sess)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create session: %v", err), http.StatusInternalServerError)
		return
	}

	middleware.Response(w, http.StatusOK, "Login successful", map[string]interface{}{
		"id":   login.UUID,
		"name": login.Name,
	})
	// Calculate and log the TTFB
	ttfb := time.Since(start)
	log.Printf("Request processed in %s (TTFB)", ttfb)
}

func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// Destroy the session
	err := middleware.DestroySession(w, r)
	if err != nil {
		middleware.Response(w, http.StatusInternalServerError, fmt.Sprintf("Failed to destroy session: %v", err), nil)
		return
	}

	// Respond with success
	middleware.Response(w, http.StatusOK, "Logged out successfully", nil)
}
