package controllers

import (
	"encoding/json"
	"net/http"
	"restaurant-booking-backend/database"
	"restaurant-booking-backend/models"

	"github.com/gorilla/mux"
)

func GetTables(w http.ResponseWriter, r *http.Request) {
	var tables []models.Table
	database.DB.Find(&tables)
	json.NewEncoder(w).Encode(tables)
}

func GetTableByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var table models.Table
	if err := database.DB.First(&table, params["id"]).Error; err != nil {
		http.Error(w, "Table not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(table)
}

func CreateTable(w http.ResponseWriter, r *http.Request) {
	var table models.Table
	if err := json.NewDecoder(r.Body).Decode(&table); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	database.DB.Create(&table)
	json.NewEncoder(w).Encode(table)
}

func UpdateTable(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var table models.Table
	if err := database.DB.First(&table, params["id"]).Error; err != nil {
		http.Error(w, "Table not found", http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&table); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	database.DB.Save(&table)
	json.NewEncoder(w).Encode(table)
}

func DeleteTable(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var table models.Table
	if err := database.DB.First(&table, params["id"]).Error; err != nil {
		http.Error(w, "Table not found", http.StatusNotFound)
		return
	}
	database.DB.Delete(&table)
	w.WriteHeader(http.StatusNoContent)
}
