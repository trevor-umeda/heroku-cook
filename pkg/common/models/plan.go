package models

import "gorm.io/gorm"

type Plan struct {
	gorm.Model        // adds ID, created_at etc.
	StartDate  string `json:"startdate`
	Sunday     string `json:"sunday"`
	Monday     string `json:"monday"`
	Tuesday    string `json:"tuesday"`

	Wednesday string `json:"wednesday"`

	Thursday string `json:"thursday"`

	Friday   string `json:"friday"`
	Saturday string `json:"saturday"`
}
