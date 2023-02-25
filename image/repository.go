package image

import (
	"GoFastApi/helper"
	"math"

	"gorm.io/gorm"
)

type Repository interface {
	Create(image Image) (Image, error)
	FindByID(id uint) (Image, error)
	Delete(image Image) error
	List(userId uint, limit int, page int, sort string) ([]Image, error)
	Exists(filename string) bool
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(image Image) (Image, error) {
	err := r.db.Create(&image).Error
	if err != nil {
		return image, err
	}
	return image, nil
}

func (r *repository) FindByID(id uint) (Image, error) {
	var image Image
	err := r.db.Where("id = ?", id).Find(&image).Error
	if err != nil {
		return image, err
	}
	return image, nil
}

func (r *repository) Delete(image Image) error {
	err := r.db.Delete(&image).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) List(userId uint, limit int, page int, sort string) ([]Image, error) {
	var images []Image

	var pagination helper.Pagination

	pagination.Limit = limit
	pagination.Page = page
	pagination.Sort = sort

	err := r.db.Where("user_id = ?", userId).Scopes(paginate(images, &pagination, r.db)).Find(&images).Error
	if err != nil {
		return images, err
	}
	return images, nil
}

func paginate(value interface{}, pagination *helper.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}

func (r *repository) Exists(filename string) bool {
	var image Image
	err := r.db.Where("image = ?", filename).Find(&image).Error
	if err != nil {
		return false
	}
	if image.ID == 0 {
		return false
	}
	return true
}
