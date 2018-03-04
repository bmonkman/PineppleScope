package models

import "time"

type Firing struct {
	ID                   uint
	StartDate            time.Time
	EndDate              time.Time
	StartDateAmbientTemp float64
	Cone                 uint
	Name                 string
	Notes                string

	TemperatureReadings []TemperatureReading
	Photos              []Photo
}

type TemperatureReading struct {
	ID          uint
	CreatedDate time.Time `gorm:"default:(datetime('now','localtime'))"`
	FiringID    uint      `gorm:"index"`
	Inner       float64
	Outer       float64
}

type Photo struct {
	ID          uint
	FiringID    uint
	CreatedDate time.Time
	photoURL    string
}