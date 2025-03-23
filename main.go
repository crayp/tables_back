package main

import (
	"log"
	"net/http"

	"restaurant-booking-backend/config"
	"restaurant-booking-backend/database"
	"restaurant-booking-backend/routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Загружаем конфигурацию
	config.LoadConfig()

	// Подключаемся к базе данных
	database.Connect()

	// Настраиваем маршруты
	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	// Настраиваем CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Разрешаем все домены, измените на список разрешенных доменов при необходимости
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(router)

	// Запускаем сервер
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", corsHandler))
}
