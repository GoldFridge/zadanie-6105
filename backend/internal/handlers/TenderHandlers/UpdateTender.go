package TenderHandlers

import (
	"backend/internal/database"
	"backend/internal/models"
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"net/http"
	"strings"
	"time"
)

// UpdateTenderRequest содержит данные для обновления тендера
type UpdateTenderRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// UpdateTender обновляет параметры существующего тендера
// @Summary Обновление существующего тендера
// @Description Обновляет параметры существующего тендера и возвращает обновленный тендер
// @Tags Tenders
// @Accept  json
// @Produce  json
// @Param tenderId path string true "ID тендера"
// @Param tender body UpdateTenderRequest true "Тело запроса для обновления тендера"
// @Success 200 {object} models.Tender "Успешное обновление тендера"
// @Failure 400 {object} map[string]string "Неверный запрос"
// @Failure 404 {object} map[string]string "Тендер не найден"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/tenders/{tenderId}/edit [patch]
func UpdateTender(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	pathSegments := strings.Split(r.URL.Path, "/")
	if len(pathSegments) < 4 {
		http.Error(w, "Invalid URL path", http.StatusBadRequest)
		return
	}

	tenderID := pathSegments[3]
	id, err := uuid.Parse(tenderID)
	if err != nil {
		http.Error(w, "Invalid tender ID", http.StatusBadRequest)
		return
	}

	var req UpdateTenderRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Поиск тендера в базе данных
	var tender models.Tender
	if err := db.First(&tender, "id = ?", id).Error; err != nil {
		if err.Error() == "record not found" {
			http.Error(w, "Tender not found", http.StatusNotFound)
		} else {
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
		return
	}

	version := models.TenderVersion{
		TenderID:    tender.ID,
		Name:        tender.Name,
		Description: tender.Description,
		Version:     tender.Version,
		CreatedAt:   time.Now(),
	}
	db.Create(&version)

	// Обновление данных тендера
	tender.Name = req.Name
	tender.Description = req.Description
	tender.UpdatedAt = time.Now()
	tender.Version++

	if err := db.Save(&tender).Error; err != nil {
		http.Error(w, "Failed to update tender", http.StatusInternalServerError)
		return
	}

	// Отправка обновленных данных тендера в ответе
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(tender)
	if err != nil {
		log.Print(err)
	}
	log.Printf("Tender with ID %s successfully updated", id)
}
