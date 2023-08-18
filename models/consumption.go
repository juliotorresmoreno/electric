package models

import "gorm.io/gorm"

type Consumption struct {
	gorm.Model

	LocationId string
	Location   *Location
}
