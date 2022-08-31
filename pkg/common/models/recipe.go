package models

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model         // adds ID, created_at etc.
	Title       string `json:"title"`
	Link        string `gorm:"unique"`
	Description string `json:"description"`
	Tags        string `json::"tags"`
}
