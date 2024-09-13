package BidHandlers

import (
	"backend/internal/database"
	"backend/internal/models"
	"encoding/json"
	"log"
	"net/http"
)

// GetMyBids возвращает список предложений текущего пользователя
// @Summary Получение списка предложений текущего пользователя
// @Description Возвращает список предложений текущего пользователя по его имени пользователя.
// @Tags Bids
// @Accept  json
// @Produce  json
// @Param username query string true "Имя пользователя" format(string)
// @Success 200 {array} models.Bid "Список предложений пользователя"
// @Failure 400 {object} map[string]string "Неверный запрос"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/bids/my [get]
func GetMyBids(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	var bids []models.Bid
	if err := db.Where("created_by = ?", username).Find(&bids).Error; err != nil {
		http.Error(w, "Failed to fetch bids", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(bids)
	if err != nil {
		log.Println("Error encoding response:", err)
	}
}
