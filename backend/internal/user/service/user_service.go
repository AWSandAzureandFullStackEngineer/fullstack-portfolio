package service

import (
	"backend/internal/user/model"
	"backend/internal/user/repository"
	"errors"
	"regexp"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) RegisterUser(user *model.User) error {
	if user.FirstName == "" || user.LastName == "" || user.Username == "" || user.Email == "" || user.Password == "" {
		return errors.New("all fields must be provided")
	}

	if !isValidEmail(user.Email) {
		return errors.New("invalid email format")
	}

	if !isStrongPassword(user.Password) {
		return errors.New("password does not meet the strength requirements")
	}

	if s.Repo.ExistsByEmail(user.Email) {
		return errors.New("a user with this email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return s.Repo.Save(user)
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

func isStrongPassword(password string) bool {
	var hasMinLen, hasDigit, hasLetter bool
	hasMinLen = len(password) >= 8

	for _, c := range password {
		if unicode.IsDigit(c) {
			hasDigit = true
		} else if unicode.IsLetter(c) {
			hasLetter = true
		}

		if hasMinLen && hasDigit && hasLetter {
			return true
		}
	}

	return hasMinLen && hasDigit && hasLetter
}
