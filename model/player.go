package model

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Ranking   int    `json:"ranking"`
}
