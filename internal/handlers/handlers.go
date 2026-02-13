package handlers

import (
	"encoding/json"
	"net/http"
	"semen_project/internal/database"
	"semen_project/models"
	"strings"
)

type Handlers struct {
	store *database.UserStore
}

func NewHandlers(store *database.UserStore) *Handlers {
	return &Handlers{store: store} // ???
}
func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}
func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	respondWithJSON(w, statusCode, map[string]string{"error": message})
}
func (h *Handlers) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.store.GetAllUsers()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Ошибка при получении пользователей")
		return
	}
	respondWithJSON(w, http.StatusOK, users)
}
func (h *Handlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Неверный формат данных")
		return
	}
	if strings.TrimSpace(user.UserName) == "" || strings.TrimSpace(user.Password) == "" {
		respondWithError(w, http.StatusBadRequest, "UserName и Password не могут быть пустыми")
		return
	}
	var userptr *models.User
	userptr, err = h.store.CreateUser(user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Ошибка при создании пользователя")
		return
	}
	respondWithJSON(w, http.StatusCreated, userptr)
}