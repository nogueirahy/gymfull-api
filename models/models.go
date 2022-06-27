package models

import (
	"time"

	"gorm.io/gorm"
)

type Member struct {
	ID        uint           `json:"member_id" gorm:"primaryKey"`
	FirstName string         `json:"firstName"`
	LastName  string         `json:"lastName"`
	CreatedAt time.Time      `gorm:"<-:create" json:"created"`
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}
