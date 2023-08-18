package models

import "gorm.io/gorm"

type City struct {
	gorm.Model

	Name string
	Code string

	CountryId uint
	Country   *Country
}
