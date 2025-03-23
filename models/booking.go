package models

import "gorm.io/gorm"

type Booking struct {
	gorm.Model
	TableID uint   `json:"table_id"`
	UserID  uint   `json:"user_id"`
	Status  string `json:"status"` // 'confirmed', 'pending', etc.
}
