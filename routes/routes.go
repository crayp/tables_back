package routes

import (
	"restaurant-booking-backend/controllers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	// GET маршруты
	router.HandleFunc("/api/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/api/restaurants", controllers.GetRestaurants).Methods("GET")
	router.HandleFunc("/api/tables", controllers.GetTables).Methods("GET")
	router.HandleFunc("/api/bookings", controllers.GetBookings).Methods("GET")

	// POST маршруты
	router.HandleFunc("/api/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/restaurants", controllers.CreateRestaurant).Methods("POST")
	router.HandleFunc("/api/tables", controllers.CreateTable).Methods("POST")
	router.HandleFunc("/api/bookings", controllers.CreateBooking).Methods("POST")
}
