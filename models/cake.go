package models

import (
	"time"

	"gorm.io/gorm"
)

type Cake struct {
	Id          int     `gorm:"primaryKey" json:"id"`
	Title       string  `gorm:"varchar(300);not null" binding:"required" json:"title" `
	Description string  `gorm:"text;not null" binding:"required" json:"description" `
	Rating      float64 `gorm:"number;not null" binding:"required" json:"rating" `
	Image       string  `gorm:"text" json:"image"`
	CreatedAt   time.Time
	UpdatedAt   time.Time `gorm:"autoUpdateTime:true"`
	DeletedAt   gorm.DeletedAt
}
