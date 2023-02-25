package users

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name          string
	Email         string
	Password      string
	EmailVerified bool
	Bio           string
	Avatar        string
	Role          string
}
