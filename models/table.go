package models

import "gorm.io/gorm"

type Table struct {
	gorm.Model
	RestaurantID uint `json:"restaurant_id"`
	Capacity     int  `json:"capacity"`
	Available    bool `json:"available"`
	Number       int  `json:"number"`
}
