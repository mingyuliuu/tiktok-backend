package models

import "time"

type Comment struct {
	CommentID      uint   `gorm:"primaryKey;autoIncrement;not null;unique;"`
	CommentContent string `gorm:"not null;"`
	CreatedAt      time.Time
}
