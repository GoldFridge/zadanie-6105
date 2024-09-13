package TenderHandlers

import (
	"backend/internal/database"
	"backend/internal/models"
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"net/http"
	"strings"
)

// RollbackTender откатывает тендер к указанной версии
// @Summary Откат тендера
// @Description Откатывает параметры тендера к указанной версии
// @Tags Tenders
// @Produce  json
// @Param tenderId path string true "ID тендера"
// @Param version path int true "Версия тендера"
// @Success 200 {object} models.Tender "Успешный откат тендера"
// @Failure 400 {object} map[string]string "Неверный запрос"
// @Failure 404 {object} map[string]string "Тендер или версия не найдены"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/tenders/{tenderId}/rollback/{version} [put]
func RollbackTender(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	pathSegments := strings.Split(r.URL.Path, "/")
	if len(pathSegments) < 5 {
		http.Error(w, "Invalid URL path", http.StatusBadRequest)
		return
	}

	tenderID := pathSegments[3]
	versionStr := pathSegments[5]

	id, err := uuid.Parse(tenderID)
	if err != nil {
		http.Error(w, "Invalid tender ID", http.StatusBadRequest)
		return
	}

	// Поиск тендера в базе данных
	var tender models.Tender
	if err := db.First(&tender, "id = ?", id).Error; err != nil {
		http.Error(w, "Tender not found", http.StatusNotFound)
		return
	}

	// Поиск указанной версии тендера
	var tenderVersion models.TenderVersion
	if err := db.First(&tenderVersion, "tender_id = ? AND version = ?", id, versionStr).Error; err != nil {
		http.Error(w, "Tender version not found", http.StatusNotFound)
		return
	}

	// Откат параметров тендера к указанной версии
	tender.Name = tenderVersion.Name
	tender.Description = tenderVersion.Description
	tender.Version = tenderVersion.Version

	if err := db.Save(&tender).Error; err != nil {
		http.Error(w, "Failed to rollback tender", http.StatusInternalServerError)
		return
	}

	// Отправка данных откатанного тендера в ответе
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(tender)
	if err != nil {
		log.Print(err)
	}
	log.Printf("Tender with ID %s successfully rolled back to version %s", id, versionStr)
}
