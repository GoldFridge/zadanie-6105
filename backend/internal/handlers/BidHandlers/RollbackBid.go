package BidHandlers

import (
	"backend/internal/database"
	"backend/internal/models"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

// RollbackBid откатывает предложение к указанной версии
// @Summary Откат предложения к указанной версии
// @Description Откатывает предложение к указанной версии и возвращает обновленные данные предложения.
// @Tags Bids
// @Accept  json
// @Produce  json
// @Param bidId path string true "ID предложения" format(uuid)
// @Param version path string true "Версия предложения"
// @Success 200 {object} models.Bid "Обновленные данные предложения"
// @Failure 400 {object} map[string]string "Неверный запрос"
// @Failure 404 {object} map[string]string "Предложение или версия не найдены"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/bids/{bidId}/rollback/{version} [put]
func RollbackBid(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	pathSegments := strings.Split(r.URL.Path, "/")
	bidID := pathSegments[3]
	version := pathSegments[5]

	var bidVersion models.BidVersion
	if err := db.First(&bidVersion, "bid_id = ? AND version = ?", bidID, version).Error; err != nil {
		http.Error(w, "Bid version not found", http.StatusNotFound)
		return
	}

	var bid models.Bid
	if err := db.First(&bid, "id = ?", bidID).Error; err != nil {
		http.Error(w, "Bid not found", http.StatusNotFound)
		return
	}

	bid.Name = bidVersion.Name
	bid.Description = bidVersion.Description
	bid.Status = bidVersion.Status
	bid.UpdatedAt = time.Now()

	if err := db.Save(&bid).Error; err != nil {
		http.Error(w, "Failed to rollback bid", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bid)
}
