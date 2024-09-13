package BidHandlers

import (
	"backend/internal/database"
	"backend/internal/models"
	"encoding/json"
	"net/http"
	"strings"
)

// GetTenderBid возвращает список предложений для указанного тендера
// @Summary Получение списка предложений для указанного тендера
// @Description Возвращает список предложений, связанных с указанным тендером.
// @Tags Bids
// @Accept  json
// @Produce  json
// @Param tenderId path string true "ID тендера" format(uuid)
// @Success 200 {array} models.Bid "Список предложений для тендера"
// @Failure 400 {object} map[string]string "Неверный запрос"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/bids/{tenderId}/list [get]
func GetTenderBid(w http.ResponseWriter, r *http.Request) {
	db := database.DB

	pathSegments := strings.Split(r.URL.Path, "/")
	if len(pathSegments) < 4 {
		http.Error(w, "Invalid URL path", http.StatusBadRequest)
		return
	}

	tenderID := pathSegments[3]

	var bids []models.Bid
	if err := db.Where("tender_id = ?", tenderID).Find(&bids).Error; err != nil {
		http.Error(w, "Failed to fetch bids for tender", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bids)
}
