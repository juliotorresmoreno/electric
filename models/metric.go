package models

import (
	"time"

	"gorm.io/gorm"
)

type Metric struct {
	gorm.Model

	LocationId         uint
	Location           *Location
	Date               time.Time
	Active             float32
	ReactiveInductive  float32
	ReactiveCapacitive float32
	Exported           float32
}
