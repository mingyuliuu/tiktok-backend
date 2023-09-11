package models

type VideoUserComment struct {
	VideoID   uint `gorm:"primaryKey;not null;"`
	UserID    uint `gorm:"primaryKey;not null;"`
	CommentID uint `gorm:"primaryKey;not null;"`

	Video   Video   `gorm:"foreignKey:VideoID"`
	User    User    `gorm:"foreignKey:UserID"`
	Comment Comment `gorm:"foreignKey:CommentID"`
}
