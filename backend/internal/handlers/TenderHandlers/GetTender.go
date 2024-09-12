package TenderHandlers

import (
	"backend/internal/database"
	"backend/internal/models"
	"encoding/json"
	"log"
	"net/http"
)

// GetTender godoc
// @Summary Получение списка тендеров
// @Description Возвращает список всех тендеров из базы данных
// @Tags Tenders
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Tender
// @Failure 500 {string} string "Error fetching data"
// @Router /api/tenders [get]
func GetTender(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	var tenders []models.Tender

	// Execute raw SQL query and scan results into tenders slice
	if err := db.Raw("SELECT * FROM tender").Scan(&tenders).Error; err != nil {
		log.Printf("Error fetching data from database: %v", err)
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}

	// Set Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Send JSON response
	if err := json.NewEncoder(w).Encode(tenders); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}
