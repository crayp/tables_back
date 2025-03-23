package main

import (
	"log"
	"net/http"

	"restaurant-booking-backend/config"
	"restaurant-booking-backend/database"
	"restaurant-booking-backend/routes"

	"github.com/gorilla/mux"
)

func main() {
	// Загружаем конфигурацию
	config.LoadConfig()

	// Подключаемся к базе данных
	database.Connect()

	// Настраиваем маршруты
	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	// Запускаем сервер
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
