package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName      string        `json:"firstName"`
	LastName       string        `json:"lastName"`
	Email          string        `json:"email"`
	GithubUsername string        `json:"githubUsername"`
	Contact        Contact       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	GithubAccount  GithubAccount `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
