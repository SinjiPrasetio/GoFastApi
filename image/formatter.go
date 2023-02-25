package image

import "time"

type ImageFormat struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"blog_id"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FormatImage(image Image) ImageFormat {
	format := ImageFormat{}
	format.ID = image.ID
	format.UserID = image.UserID
	format.Image = image.Image
	format.CreatedAt = image.CreatedAt
	format.UpdatedAt = image.UpdatedAt

	return format
}
