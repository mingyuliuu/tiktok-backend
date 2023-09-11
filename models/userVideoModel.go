package models

type UserVideo struct {
	UserID  uint `gorm:"primaryKey;not null;"`
	VideoID uint `gorm:"primaryKey;not null;"`

	User  User  `gorm:"foreignKey:UserID"`
	Video Video `gorm:"foreignKey:VideoID"`
}
