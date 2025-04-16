package models

import (
	"time"

	"gorm.io/gorm"
)

type Emotion struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Mood      string         `json:"mood"` // happy / sad / angry / etc
	Note      string         `json:"note"` // optional
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
