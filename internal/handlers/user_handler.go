package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"myapp/internal/service"
	"myapp/pkg/logger"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	Service *service.UserService
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	user, err := h.Service.GetUser(r.Context(), id)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "encoding error", http.StatusInternalServerError)
		return
	}

	logger.Infof("User fetched: %d", id)
}
