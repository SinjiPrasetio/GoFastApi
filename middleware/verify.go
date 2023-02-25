package middleware

import (
	"GoFastApi/helper"
	"GoFastApi/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Verify(userService users.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("currentUser").(users.User)
		if !user.EmailVerified {
			response := helper.ResponseFormatter("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
	}
}
