package models

import "gorm.io/gorm"

type Country struct {
	gorm.Model

	Name string
	Code string

	CountryId uint
}
