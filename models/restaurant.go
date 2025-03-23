package models

import "gorm.io/gorm"

type Restaurant struct {
	gorm.Model
	Name        string  `json:"name"`
	Location    string  `json:"location"`
	Description string  `json:"description"`
	OwnerID     uint    `json:"owner_id"`
	Tables      []Table `json:"tables"`
}
