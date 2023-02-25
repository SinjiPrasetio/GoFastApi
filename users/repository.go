package users

import "gorm.io/gorm"

// Deklarasi Repository untuk fungsi Database
type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByID(ID uint) (User, error)
	UpdateBioData(ID uint, text string) (User, error)
	Update(user User) (User, error)
	AvatarExists(filename string) bool
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByID(ID uint) (User, error) {
	var user User
	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) UpdateBioData(ID uint, text string) (User, error) {
	user, err := r.FindByID(ID)
	if err != nil {
		return user, err
	}

	r.db.First(&user)
	user.Bio = text
	r.db.Save(&user)

	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) AvatarExists(filename string) bool {
	var user User
	err := r.db.Where("avatar = ?", filename).Find(&user).Error
	if err != nil {
		return true
	}
	if user.ID != 0 {
		return true
	}
	return false
}
