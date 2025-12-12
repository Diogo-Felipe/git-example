package model

import "github.com/google/uuid"

type Player struct {
	ID                 string    `gorm:"primaryKey"`
	Url                string    `gorm:"index"`
}