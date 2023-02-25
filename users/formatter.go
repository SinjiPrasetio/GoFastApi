package users

import "time"

type UserFormatter struct {
	ID            uint      `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Token         string    `json:"token"`
	EmailVerified bool      `json:"email_verified"`
	Bio           string    `json:"bio"`
	Avatar        string    `json:"avatar"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type UserNoTokenFormatter struct {
	ID            uint      `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	EmailVerified bool      `json:"email_verified"`
	Bio           string    `json:"bio"`
	Avatar        string    `json:"avatar"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type EmailFormatter struct {
	IsEmailAvailable bool `json:"is_email_available"`
}

type TestData struct {
	Id   uint
	Name string
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:            user.ID,
		Name:          user.Name,
		Email:         user.Email,
		Token:         token,
		EmailVerified: user.EmailVerified,
		Bio:           user.Bio,
		Avatar:        user.Avatar,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}

	return formatter
}

func FormatEmail(isAvail bool) EmailFormatter {
	formatter := EmailFormatter{
		IsEmailAvailable: isAvail,
	}

	return formatter
}

func FormatUserWithoutToken(user User) UserNoTokenFormatter {
	formatter := UserNoTokenFormatter{
		ID:            user.ID,
		Name:          user.Name,
		Email:         user.Email,
		EmailVerified: user.EmailVerified,
		Bio:           user.Bio,
		Avatar:        user.Avatar,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}

	return formatter
}
