package handlers

import (
	"net/http"
	"to-do-list/internal/domain/service"
	"to-do-list/internal/infrastructure/middleware"
)

type TaskHandler struct {
	service *service.TaskService
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	authenticated, err := middleware.IsAuthenticated(r)
	if err != nil || !authenticated {
		middleware.Response(w, http.StatusUnauthorized, "Unauthorized access", nil)
		return
	}
}
