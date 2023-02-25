package middleware

type TokenValidate struct {
	Token string `json:"token" binding:"required"`
}
