package models

import "time"

type Post struct {
	ID              uint   `gorm:"primaryKey"`
	Baslik          string `gorm:"not null"`
	İcerik          string `gorm:"type:text"`
	ResimYolu       string
	OlusturmaTarihi time.Time
}
