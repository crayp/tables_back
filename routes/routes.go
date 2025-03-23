package routes

import (
	"restaurant-booking-backend/controllers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/api/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/api/restaurants", controllers.GetRestaurants).Methods("GET")
	router.HandleFunc("/api/tables", controllers.GetTables).Methods("GET")
	router.HandleFunc("/api/bookings", controllers.GetBookings).Methods("GET")
}
