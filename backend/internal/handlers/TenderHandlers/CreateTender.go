package TenderHandlers

import (
	"backend/internal/database"
	"backend/internal/models"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type CreateTenderRequest struct {
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	ServiceType     string    `json:"serviceType"`
	OrganizationID  uuid.UUID `json:"organizationId"`
	CreatorUsername string    `json:"creatorUsername"`
}

// CreateTender создает новый тендер
// @Summary Создание нового тендера
// @Description Создает новый тендер и возвращает его
// @Tags Tenders
// @Accept  json
// @Produce  json
// @Param tender body models.Tender true "Тело запроса для создания тендера"
// @Success 200 {object} models.Tender "Успешное создание тендера"
// @Router /api/tenders/new [post]
func CreateTender(w http.ResponseWriter, r *http.Request) {
	db := database.DB

	var req CreateTenderRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	newTender := models.Tender{
		ID:              uuid.New(),
		Name:            req.Name,
		Description:     req.Description,
		OrganizationID:  req.OrganizationID,
		CreatorUsername: req.CreatorUsername,
		CreatedBy:       uuid.New(),
		Status:          models.TenderCreated,
		Version:         1,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	if err := db.Create(&newTender).Error; err != nil {
		http.Error(w, "Failed to create tender", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(newTender)
	if err != nil {
		log.Print(err)
	}
}
