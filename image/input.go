package image

type DeleteInput struct {
	ID uint `json:"id" binding:"required"`
}

type ListInput struct {
	Limit  int    `json:"limit" binding:"required"`
	Page   int    `json:"page" binding:"required"`
	Sort   string `json:"sort" binding:"required"`
	Search string `json:"search"`
}
