package users

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	GetUserByID(ID uint) (User, error)
	UpdateBio(ID uint, text string) (User, error)
	UploadAvatar(ID uint, url string) (User, error)
	AvatarExists(filename string) bool
	RemoveAvatar(ID uint) bool
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Role = "user"
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)

	saved, err := s.repository.Save(user)
	if err != nil {
		return user, err
	}
	return saved, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User not found.")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) GetUserByID(ID uint) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) UpdateBio(ID uint, text string) (User, error) {
	user, err := s.repository.UpdateBioData(ID, text)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) UploadAvatar(ID uint, fileLocation string) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	user.Avatar = fileLocation

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) AvatarExists(filename string) bool {
	return s.repository.AvatarExists(filename)
}

func (s *service) RemoveAvatar(ID uint) bool {
	user, err := s.GetUserByID(ID)
	if err != nil {
		return false
	}
	user.Avatar = ""

	_, err = s.repository.Update(user)
	if err != nil {
		return false
	}
	return true
}
