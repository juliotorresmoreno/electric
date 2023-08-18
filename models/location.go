package models

import "gorm.io/gorm"

type Location struct {
	gorm.Model

	OwnerId     uint
	Owner       *User
	Address     string
	Description string
	Latitude    float32
	Longitude   float32
}
