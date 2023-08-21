package model

// References: https://gorm.io/docs/models.html

type Users struct {
	UserID       int    `gorm:"not null;autoIncrement=1;primaryKey"`
	UserName     string `gorm:"not null;"`
	UserPassword string `gorm:"not null;"`
}
