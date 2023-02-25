package image

import (
	"GoFastApi/users"
	"time"
)

type Image struct {
	ID        uint
	UserID    uint
	User      users.User
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
