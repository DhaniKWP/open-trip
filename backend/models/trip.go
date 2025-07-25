package models

import "gorm.io/gorm"

type Trip struct {
	gorm.Model
	Name        string  `json:"name"`
	Location    string  `json:"location"`
	DateStart   string  `json:"date_start"`
	DateEnd     string  `json:"date_end"`
	Capacity    int     `json:"capacity"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
}
