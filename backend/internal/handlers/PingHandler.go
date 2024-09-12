package handlers

import (
	"fmt"
	"log"
	"net/http"
)

// PingHandler @Summary Проверка статуса сервера
// @Description Возвращает "ok" с кодом состояния 200
// @Tags Health
// @Accept  json
// @Produce  json
// @Success 200 {string} string "ok"
// @Router /api/ping [get]
func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprint(w, "ok")
	if err != nil {
		log.Fatal(err)
	}
}
