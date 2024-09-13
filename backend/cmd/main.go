package main

import (
	_ "backend/docs"
	"backend/internal/database"
	"backend/internal/handlers"
	"backend/internal/handlers/BidHandlers"
	"backend/internal/handlers/TenderHandlers"
	_ "backend/internal/models"
	"fmt"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

// @title TenderAPI
// @version 1.0
// @description API Server for TodoList Application

// @host cnrprod1725725290-team-77871-32639.avito2024.codenrock.com
// @BasePath /

func main() {
	database.InitDatabase()
	http.HandleFunc("/api/ping", handlers.PingHandler)
	http.HandleFunc("/api/tenders", TenderHandlers.GetTender)
	http.HandleFunc("/api/tenders/new", TenderHandlers.CreateTender)
	http.HandleFunc("/api/tenders/{tenderId}/edit", TenderHandlers.UpdateTender)
	http.HandleFunc("/api/tenders/{tenderId}/rollback/{version}", TenderHandlers.RollbackTender)
	http.HandleFunc("/api/bids/new", BidHandlers.CreateBid)
	http.HandleFunc("/api/bids/my", BidHandlers.GetMyBids)
	http.HandleFunc("/api/bids/{tenderId}/list", BidHandlers.GetMyBids)
	http.HandleFunc("/api/bids/{bidId}/edit", BidHandlers.UpdateBid)
	http.HandleFunc("/api/bids/{bidId}/rollback/{version}", BidHandlers.RollbackBid)
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	fmt.Println("Сервер запущен на http://localhost:8080/api")
	fmt.Println("Swagger: http://localhost:8080/swagger/index.html")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error start server")
	}
}
