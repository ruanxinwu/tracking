package model

import "github.com/jinzhu/gorm"

type JwtBlackList struct {
	gorm.Model
	Jwt string `gorm:"type:text;comment:'jwt'"`
}
