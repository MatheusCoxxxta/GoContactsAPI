package models

import "gorm.io/gorm"

type GithubAccount struct {
	gorm.Model
	AvatarUrl    string `json:"avatar_url"`
	Public_Repos uint64 `json:"public_repos"`
	UserID       uint64 `json:"user_id"`
}
