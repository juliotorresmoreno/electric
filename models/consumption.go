package models

import (
	"gorm.io/gorm"
)

type Consumption struct {
	gorm.Model

	Address            string
	Period             string
	Active             float32
	ReactiveInductive  float32
	ReactiveCapacitive float32
	Exported           float32
}
