// mocks/user_repository_mock.go
package mocks

import (
	"clean-architecture/domain"

	"github.com/stretchr/testify/mock"
)

// UserRepositoryMock is a manually created mock for the UserRepository interface
type UserRepositoryMock struct {
	mock.Mock
}

// Register is a mock implementation of the UserRepository.Register method
func (m *UserRepositoryMock) Register(user domain.User) (domain.User, error) {
	args := m.Called(user)
	return args.Get(0).(domain.User), args.Error(1)
}

// FindByUsername is a mock implementation of the UserRepository.FindByUsername method
func (m *UserRepositoryMock) FindByUsername(username string) (domain.User, error) {
	args := m.Called(username)
	return args.Get(0).(domain.User), args.Error(1)
}

// GetUsers is a mock implementation of the UserRepository.GetUsers method
func (m *UserRepositoryMock) GetUsers() ([]domain.User, error) {
	args := m.Called()
	return args.Get(0).([]domain.User), args.Error(1)
}
