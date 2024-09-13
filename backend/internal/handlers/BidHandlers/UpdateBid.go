package BidHandlers

import (
	"backend/internal/database"
	"backend/internal/models"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type UpdateBidRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// UpdateBid обновляет существующее предложение
// @Summary Обновление существующего предложения
// @Description Обновляет параметры предложения и возвращает обновленные данные предложения.
// @Tags Bids
// @Accept  json
// @Produce  json
// @Param bidId path string true "ID предложения" format(uuid)
// @Param bid body UpdateBidRequest true "Тело запроса для обновления предложения"
// @Success 200 {object} models.Bid "Успешное обновление предложения"
// @Failure 400 {object} map[string]string "Неверный запрос"
// @Failure 404 {object} map[string]string "Предложение не найдено"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/bids/{bidId}/edit [patch]
func UpdateBid(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	pathSegments := strings.Split(r.URL.Path, "/")
	if len(pathSegments) < 4 {
		http.Error(w, "Invalid URL path", http.StatusBadRequest)
		return
	}

	// Извлекаем bidId из URL
	bidID := pathSegments[3]
	var req UpdateBidRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var bid models.Bid
	if err := db.First(&bid, "id = ?", bidID).Error; err != nil {
		http.Error(w, "Bid not found", http.StatusNotFound)
		return
	}

	bid.Name = req.Name
	bid.Description = req.Description
	bid.UpdatedAt = time.Now()

	if err := db.Save(&bid).Error; err != nil {
		http.Error(w, "Failed to update bid", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bid)
}
