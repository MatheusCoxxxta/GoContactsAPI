package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Phone   string `json:"phone"`
	Twitter string `json:"twitter"`
	UserID  uint64 `json:"user_id"`
}
