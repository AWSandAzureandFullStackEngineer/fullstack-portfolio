package user

import (
	"backend/internal/user/model"
	"backend/internal/user/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

type MockRepository struct {
	users map[string]*model.User
}

func NewMockRepository() *MockRepository {
	return &MockRepository{users: make(map[string]*model.User)}
}

func (m *MockRepository) Save(user *model.User) error {
	if _, exists := m.users[user.Email]; exists {
		return errors.New("user already exists")
	}
	m.users[user.Email] = user
	return nil
}

func (m *MockRepository) ExistsByEmail(email string) bool {
	_, exists := m.users[email]
	return exists
}

func TestRegisterUser(t *testing.T) {
	repo := NewMockRepository()
	userService := service.NewUserService(repo)

	user := &model.User{
				FirstName: "New",
					LastName:  "User",
					Username:  "newuser",
					Email:     "newuser@example.com",
					Password:  "Password1",
				
	}

	err := userService.RegisterUser(user)
	if err != nil {
					t.Errorf("RegisterUser() error = %v, wantErr %v", err, nil)
	}

	savedUser, exists := repo.users[user.Email]
	if !exists {
					t.Fatal("User was not saved in the repository")
	}

	err = bcrypt.CompareHashAndPassword([]byte(savedUser.Password), []byte("Password1"))
	if err != nil {
					t.Errorf("The password should be hashed and match the original, err: %v", err)
	}
}

func TestRegisterUserWeakPassword(t *testing.T) {
	repo := NewMockRepository()
	userService := service.NewUserService(repo)

	user := &model.User{
		FirstName: "New",
		LastName:  "User",
		Username:  "newuser",
		Email:     "newuser@example.com",
		Password:  "1234567", 
	}

	err := userService.RegisterUser(user)
	assert.Error(t, err, "Should fail due to weak password")
	assert.Contains(t, err.Error(), "password does not meet the strength requirements", "Error message should indicate the password strength issue")
}

func TestRegisterUserInvalidEmail(t *testing.T) {
	repo := NewMockRepository()
	userService := service.NewUserService(repo)

	user := &model.User{
		FirstName: "New",
		LastName:  "User",
		Username: "newuser",
		Email:    "not-an-email",
		Password: "securepassword",
	}

	err := userService.RegisterUser(user)
	assert.Error(t, err, "Should fail due to invalid email format")
}
