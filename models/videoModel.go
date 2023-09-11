package models

import "time"

type Video struct {
	VideoID    uint   `gorm:"primaryKey;autoIncrement;not null;unique;"`
	VideoURL   string `gorm:"not null;"`
	VideoCover string `gorm:"not null;"`
	VideoTitle string `gorm:"not null;default:say something;"`
	CreatedAt  time.Time
}
