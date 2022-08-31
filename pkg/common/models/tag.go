package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model        // adds ID, created_at etc.
	Tag        string `json:"tag"`
	Recipes    string `json:"recipes"`
}
