package service

import (
	"atmail/internal/helper"
	"atmail/internal/model"
	"atmail/internal/repository"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type userService struct {
	userRepository repository.UserRepository
}

type UserService interface {
	Delete(id uint) error
	Get(id uint) (*model.User, int, error)
	GetAll() (*[]model.User, error)
	Save(req model.UserRequest) (resp *model.User, err error)
	Update(req model.User) (*model.User, error)
	ValidateNewUser(req model.UserRequest) error
	ValidateExistingUser(req model.User) (int, error)
	ValidateID(id uint) (int, error)
}

func NewUserService(repository repository.UserRepository) UserService {
	service := new(userService)
	service.userRepository = repository
	return service
}

// Get user by ID
func (u *userService) Get(id uint) (*model.User, int, error) {
	user, err := u.userRepository.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, http.StatusNotFound, errors.New("user not found")
		}
		return nil, http.StatusBadRequest, err
	}

	return user, http.StatusOK, nil
}

// Get all users
func (u *userService) GetAll() (*[]model.User, error) {
	users, err := u.userRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Create new user
func (u *userService) Save(req model.UserRequest) (user *model.User, err error) {
	var r repository.User
	r.Username = req.Username
	r.Email = req.Email
	r.Age = req.Age

	updated, err := u.userRepository.Save(r)
	if err != nil {
		return nil, err
	}
	return updated, nil
}

// Validate requests for new users
func (u *userService) ValidateNewUser(req model.UserRequest) error {
	var id *uint
	if err := u.validateEmail(req.Email, id); err != nil {
		return err
	}
	if err := u.validateUsername(req.Username, id); err != nil {
		return err
	}
	if !helper.IsAgeValid(req.Age) {
		return errors.New("invalid age")
	}
	return nil
}

// Validate requests for existing users
func (u *userService) ValidateExistingUser(req model.User) (int, error) {
	statusCode, err := u.validateID(req.ID)
	if err != nil {
		return statusCode, err
	}
	if err := u.validateEmail(req.Email, &req.ID); err != nil {
		return http.StatusBadRequest, err
	}
	if err := u.validateUsername(req.Username, &req.ID); err != nil {
		return http.StatusBadRequest, err
	}
	if !helper.IsAgeValid(req.Age) {
		return http.StatusBadRequest, errors.New("invalid age")
	}
	return http.StatusOK, nil
}

// validate if ID exists in the database
func (u *userService) validateID(id uint) (int, error) {
	_, err := u.userRepository.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusNotFound, errors.New("no record found")
		}
		return http.StatusBadRequest, err
	}
	return http.StatusOK, nil
}

// validate if email exists in the database
func (u *userService) validateEmail(email string, id *uint) error {
	if len(email) == 0 {
		return errors.New("email is required")
	}
	if !helper.IsEmailValid(email) {
		return errors.New("invalid email")
	}

	isUnique, err := u.userRepository.IsEmailUnique(id, email)
	if err != nil {
		return err
	}
	if !isUnique {
		return errors.New("email already exists")
	}

	return nil
}

// validate if username exists in the database
func (u *userService) validateUsername(username string, id *uint) error {
	if len(username) == 0 {
		return errors.New("username is required")
	}
	if !helper.IsUsernameValid(username) {
		return errors.New("invalid username")
	}

	isUnique, err := u.userRepository.IsUsernameUnique(id, username)
	if err != nil {
		return err
	}
	if !isUnique {
		return errors.New("username already exists")
	}

	return nil
}

// Update changes
func (u *userService) Update(req model.User) (*model.User, error) {
	user, err := u.userRepository.GetUser(req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no record found")
		}
		return nil, err
	}
	user.Username = req.Username
	user.Email = req.Email
	user.Age = req.Age
	updated, err := u.userRepository.Update(*user)
	if err != nil {
		return nil, err
	}
	return updated, nil
}

// Delete a user
func (u *userService) Delete(id uint) error {
	if err := u.userRepository.Delete(id); err != nil {
		return err
	}
	return nil
}

// Validate if ID exists in the DB
func (u *userService) ValidateID(id uint) (int, error) {
	return u.validateID(id)
}
