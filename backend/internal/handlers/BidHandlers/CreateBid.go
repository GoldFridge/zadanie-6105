package BidHandlers

import (
	"backend/internal/database"
	"backend/internal/models"
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

type CreateBidRequest struct {
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Status          string    `json:"status"`
	TenderID        uuid.UUID `json:"tenderId"`
	OrganizationID  uuid.UUID `json:"organizationId"`
	CreatorUsername string    `json:"creatorUsername"`
}

// CreateBid создает новое предложение
// @Summary Создание нового предложения
// @Description Создает новое предложение для существующего тендера
// @Tags Bids
// @Accept json
// @Produce json
// @Param request body CreateBidRequest true "Данные для создания нового предложения"
// @Success 200 {object} models.Bid "Успешное создание предложения"
// @Failure 400 {object} map[string]string "Неверный запрос"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/bids/new [post]
func CreateBid(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	var req CreateBidRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	bid := models.Bid{
		ID:          uuid.New(),
		Name:        req.Name,
		Description: req.Description,
		Status:      models.BidStatus(req.Status),
		TenderID:    req.TenderID,
		CreatedBy:   uuid.New(), // можно заменить логикой поиска пользователя по username
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := db.Create(&bid).Error; err != nil {
		http.Error(w, "Failed to create bid", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(bid)
	if err != nil {
		log.Println(err)
	}
}
