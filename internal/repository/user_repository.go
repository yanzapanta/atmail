package repository

import (
	"atmail/internal/config"
	"atmail/internal/model"
	"errors"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type userRepository struct {
}

type UserRepository interface {
	Delete(id uint) error
	Get(id uint) (*model.User, error)
	GetAll() (*[]model.User, error)
	GetUser(id uint) (*User, error)
	IsEmailUnique(id *uint, email string) (bool, error)
	IsUsernameUnique(id *uint, email string) (bool, error)
	Save(user User) (*model.User, error)
	Update(user User) (*model.User, error)
}

func NewUserRepository() UserRepository {
	repo := new(userRepository)
	return repo
}

func (u *userRepository) Get(id uint) (*model.User, error) {
	var user User
	user.ID = id
	if err := config.DB().Take(&user).Error; err != nil {
		return nil, err
	}
	var m model.User
	copier.Copy(&m, user)
	return &m, nil
}

func (u *userRepository) GetAll() (*[]model.User, error) {
	var users []User
	if err := config.DB().Find(&users).Error; err != nil {
		return nil, err
	}
	var m []model.User
	copier.Copy(&m, users)
	return &m, nil
}

func (u *userRepository) Save(user User) (*model.User, error) {
	if err := config.DB().Create(&user).Error; err != nil {
		return nil, err
	}
	var m model.User
	copier.Copy(&m, user)
	return &m, nil
}

func (u *userRepository) IsEmailUnique(id *uint, email string) (bool, error) {
	var user User
	query := config.DB().Where("email = ?", email)
	if id != nil {
		query = query.Where("id != ?", *id)
	}

	if err := query.First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func (u *userRepository) IsUsernameUnique(id *uint, username string) (bool, error) {
	query := config.DB().Where("username = ?", username)
	if id != nil {
		query = query.Where("id != ?", *id)
	}
	if err := query.First(&User{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func (u *userRepository) GetUser(id uint) (*User, error) {
	var user User
	user.ID = id
	if err := config.DB().First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) Update(user User) (*model.User, error) {
	if err := config.DB().Save(&user).Error; err != nil {
		return nil, err
	}
	var m model.User
	copier.Copy(&m, user)
	return &m, nil
}

func (u *userRepository) Delete(id uint) error {
	var user User
	user.ID = id
	if err := config.DB().Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
