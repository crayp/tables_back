package controllers

import (
	"encoding/json"
	"net/http"
	"restaurant-booking-backend/database"
	"restaurant-booking-backend/models"

	"github.com/gorilla/mux"
)

func GetBookings(w http.ResponseWriter, r *http.Request) {
	var bookings []models.Booking
	database.DB.Find(&bookings)
	json.NewEncoder(w).Encode(bookings)
}

func GetBookingByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var booking models.Booking
	if err := database.DB.First(&booking, params["id"]).Error; err != nil {
		http.Error(w, "Booking not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(booking)
}

func CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking models.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	database.DB.Create(&booking)
	json.NewEncoder(w).Encode(booking)
}

func UpdateBooking(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var booking models.Booking
	if err := database.DB.First(&booking, params["id"]).Error; err != nil {
		http.Error(w, "Booking not found", http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	database.DB.Save(&booking)
	json.NewEncoder(w).Encode(booking)
}

func DeleteBooking(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var booking models.Booking
	if err := database.DB.First(&booking, params["id"]).Error; err != nil {
		http.Error(w, "Booking not found", http.StatusNotFound)
		return
	}
	database.DB.Delete(&booking)
	w.WriteHeader(http.StatusNoContent)
}
