package model

import (
	// "github.com/google/uuid"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model // Adds some metadata fields to the table
	// ID         uint `gorm:"primaryKey"` //  uuid.UUID `gorm:"type:uuid"` // Explicitly specify the type to be uuid
	Title    string
	SubTitle string
	Artikel  string
}
