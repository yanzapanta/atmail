package helper

import (
	"errors"
	"regexp"
	"strconv"
)

func IsEmailValid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

func IsUsernameValid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._\-]{3,}$`)
	return emailRegex.MatchString(email)
}

func IsAgeValid(a int) bool {
	return a > 0 && a < 100
}

func CleanID(id string) (*uint, error) {
	newID, err := strconv.ParseUint(id, 10, 64)
	if err != nil || newID == 0 {
		return nil, errors.New("invalid ID")
	}

	cleanID := uint(newID)
	return &cleanID, nil
}
