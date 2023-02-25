package middleware

import (
	"GoFastApi/core"
	"GoFastApi/helper"
	"GoFastApi/users"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Auth(authService core.AuthService, userService users.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.ResponseFormatter("Unauthorized", http.StatusUnauthorized, "error", "Access Denied : You're not authorized to call this API!")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Split Bearer dan Token
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.ResponseFormatter("Unauthorized", http.StatusUnauthorized, "error", "Fail to validate token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.ResponseFormatter("Unauthorized", http.StatusUnauthorized, "error", "Token is not valid")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := uint(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.ResponseFormatter("Unauthorized", http.StatusUnauthorized, "error", "User is unauthorized")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentUser", user)
	}
}

func Validate(authService core.AuthService, userService users.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			return
		}

		// Split Bearer dan Token
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			return
		}

		userID := uint(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			return
		}

		c.Set("currentUser", user)
	}
}
