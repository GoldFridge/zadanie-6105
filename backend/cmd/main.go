package main

import (
	"backend/internal/database"
	"backend/internal/handlers"
	"backend/internal/handlers/TenderHandlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	database.InitDatabase()
	http.HandleFunc("/api/ping", handlers.PingHandler)
	http.HandleFunc("/api/tenders", TenderHandlers.GetTender)
	fmt.Println("Сервер запущен на http://localhost:8080/api")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error start server")
	}
}
