package handler

import (
	"GoFastApi/core"
	"GoFastApi/helper"
	"GoFastApi/users"
	"errors"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService users.Service
	authService core.AuthService
}

func NewUserHandler(userService users.Service, authService core.AuthService) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input users.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ValidationErrorFormatter(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.ResponseFormatter("Failed to register", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, err := h.userService.RegisterUser(input)
	if err != nil {

		response := helper.ResponseFormatter("Failed to register", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(user.ID)
	if err != nil {
		response := helper.ResponseFormatter("Failed to register", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := users.FormatUser(user, token)

	response := helper.ResponseFormatter("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) LoginUser(c *gin.Context) {
	var input users.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ValidationErrorFormatter(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.ResponseFormatter("Failed to login", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, err := h.userService.Login(input)
	if err != nil {
		response := helper.ResponseFormatter("Failed to login", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(user.ID)
	if err != nil {
		response := helper.ResponseFormatter("Failed to login", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := users.FormatUser(user, token)

	response := helper.ResponseFormatter("Login success", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Validate(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(users.User)
	token, err := h.authService.GenerateToken(currentUser.ID)
	if err != nil {
		response := helper.ResponseFormatter("Failed to Validate", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	user := users.FormatUser(currentUser, token)

	response := helper.ResponseFormatter("Validate success", http.StatusOK, "success", user)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmail(c *gin.Context) {
	var input users.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ValidationErrorFormatter(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.ResponseFormatter("Failed to validate email", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isAvail, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errors := helper.ValidationErrorFormatter(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.ResponseFormatter("Failed to validate email", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := users.FormatEmail(isAvail)

	response := helper.ResponseFormatter("Check Email success", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) BioUpdateHandler(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(users.User)
	userID := currentUser.ID
	var input users.BioUpdateInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ValidationErrorFormatter(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.ResponseFormatter("Failed to update bio", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, err := h.userService.UpdateBio(userID, input.Text)
	if err != nil {
		response := helper.ResponseFormatter("Failed to update bio", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := users.FormatUserWithoutToken(user)

	response := helper.ResponseFormatter("Update bio successs", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(users.User)
	userID := currentUser.ID
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ResponseFormatter("Failed to upload avatar", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
	}
	if currentUser.Avatar != "" {
		if _, err := os.Stat("public/img/avatar/" + currentUser.Avatar); !errors.Is(err, os.ErrNotExist) {
			err = os.Remove("public/img/avatar/" + currentUser.Avatar)
			if err != nil {
				data := gin.H{"is_uploaded": false}
				response := helper.ResponseFormatter("Failed to upload avatar", http.StatusBadRequest, "error", data)
				c.JSON(http.StatusBadRequest, response)
				return
			}
			removeAvatar := h.userService.RemoveAvatar(userID)
			if !removeAvatar {
				data := gin.H{"is_uploaded": false}
				response := helper.ResponseFormatter("Failed to upload avatar", http.StatusBadRequest, "error", data)
				c.JSON(http.StatusBadRequest, response)
				return
			}
		}
	}

	condition := true
	var filename string
	for ok := true; ok; ok = condition {
		filename = helper.RandomString(30) + filepath.Ext(file.Filename)
		condition = h.userService.AvatarExists(filename)
	}

	path := "public/img/avatar/" + filename

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ResponseFormatter("Failed to upload avatar", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.UploadAvatar(userID, filename)
	if err != nil {
		err = os.Remove(path)
		if err != nil {
			data := gin.H{"is_uploaded": false}
			response := helper.ResponseFormatter("Failed to upload avatar", http.StatusBadRequest, "error", data)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		data := gin.H{"is_uploaded": false}
		response := helper.ResponseFormatter("Failed to upload avatar", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ResponseFormatter("Succes uploaded avatar", http.StatusOK, "success", path)
	c.JSON(http.StatusOK, response)
}
