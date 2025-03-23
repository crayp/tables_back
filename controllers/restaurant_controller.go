package controllers

import (
	"encoding/json"
	"net/http"
	"restaurant-booking-backend/database"
	"restaurant-booking-backend/models"

	"github.com/gorilla/mux"
)

func GetRestaurants(w http.ResponseWriter, r *http.Request) {
	var restaurants []models.Restaurant
	database.DB.Preload("Tables").Find(&restaurants)
	json.NewEncoder(w).Encode(restaurants)
}

func GetRestaurantByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var restaurant models.Restaurant
	if err := database.DB.Preload("Tables").First(&restaurant, params["id"]).Error; err != nil {
		http.Error(w, "Restaurant not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(restaurant)
}

func CreateRestaurant(w http.ResponseWriter, r *http.Request) {
	var restaurant models.Restaurant
	if err := json.NewDecoder(r.Body).Decode(&restaurant); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	database.DB.Create(&restaurant)
	json.NewEncoder(w).Encode(restaurant)
}

func UpdateRestaurant(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var restaurant models.Restaurant
	if err := database.DB.First(&restaurant, params["id"]).Error; err != nil {
		http.Error(w, "Restaurant not found", http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&restaurant); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	database.DB.Save(&restaurant)
	json.NewEncoder(w).Encode(restaurant)
}

func DeleteRestaurant(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var restaurant models.Restaurant
	if err := database.DB.First(&restaurant, params["id"]).Error; err != nil {
		http.Error(w, "Restaurant not found", http.StatusNotFound)
		return
	}
	database.DB.Delete(&restaurant)
	w.WriteHeader(http.StatusNoContent)
}
