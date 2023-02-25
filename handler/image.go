package handler

import (
	"GoFastApi/core"
	"GoFastApi/helper"
	"GoFastApi/image"
	"GoFastApi/users"
	"errors"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type imageHandler struct {
	imageService image.Service
	authService  core.AuthService
}

func NewImageHandler(imageService image.Service, authService core.AuthService) *imageHandler {
	return &imageHandler{imageService, authService}
}

func (i *imageHandler) Upload(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(users.User)
	userID := currentUser.ID
	file, err := c.FormFile("thumbnail")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ResponseFormatter("Failed to upload thumbnail", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	condition := true
	var filename string
	for ok := true; ok; ok = condition {
		filename = helper.RandomString(30) + filepath.Ext(file.Filename)
		condition = i.imageService.Exists(filename)
	}

	images, err := i.imageService.Upload(userID, filename)
	if err != nil {
		data := image.FormatImage(images)
		response := helper.ResponseFormatter("Failed to upload thumbnail", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	path := "public/img/thumbnail/" + filename

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := image.FormatImage(images)
		response := helper.ResponseFormatter("Failed to upload thumbnail", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := image.FormatImage(images)
	response := helper.ResponseFormatter("Succes uploaded thumbnail", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (i *imageHandler) Delete(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(users.User)
	userID := currentUser.ID
	var input image.DeleteInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ValidationErrorFormatter(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.ResponseFormatter("Parameter not given", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	image, err := i.imageService.FindByID(input.ID)

	if err != nil {
		response := helper.ResponseFormatter("Failed to find image", http.StatusUnprocessableEntity, "error", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if image.UserID != userID {
		response := helper.ResponseFormatter("Failed to find image", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if _, err := os.Stat("public/img/thumbnail/" + image.Image); !errors.Is(err, os.ErrNotExist) {
		err = os.Remove("public/img/thumbnail/" + image.Image)
		if err != nil {
			data := gin.H{"is_uploaded": false}
			response := helper.ResponseFormatter("Failed to delete thumbnail", http.StatusBadRequest, "error", data)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		err := i.imageService.Delete(image)
		if err != nil {
			data := gin.H{"is_uploaded": false}
			response := helper.ResponseFormatter("Failed to delete thumbnail", http.StatusBadRequest, "error", data)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	response := helper.ResponseFormatter("Succes delete image", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
