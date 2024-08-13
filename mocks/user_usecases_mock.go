// mocks/user_usecase_mock.go
package mocks

import (
	"clean-architecture/domain"

	"github.com/stretchr/testify/mock"
)

// UserUsecaseMock is a mock implementation of the UserUsecase interface.
type UserUsecaseMock struct {
	mock.Mock
}

// Register mocks the Register method.
func (m *UserUsecaseMock) Register(user domain.User) (domain.User, error) {
	args := m.Called(user)
	return args.Get(0).(domain.User), args.Error(1)
}

// Login mocks the Login method.
func (m *UserUsecaseMock) Login(username, password string) (string, error) {
	args := m.Called(username, password)
	return args.String(0), args.Error(1)
}

// GetUsers mocks the GetUsers method.
func (m *UserUsecaseMock) GetUsers() ([]domain.User, error) {
	args := m.Called()
	return args.Get(0).([]domain.User), args.Error(1)
}
