// mocks/password_service_mock.go
package mocks

import (
	"github.com/stretchr/testify/mock"
)

// PasswordServiceMock is a mock implementation of the PasswordService interface.
type PasswordServiceMock struct {
	mock.Mock
}

// HashPassword mocks the HashPassword method.
func (m *PasswordServiceMock) HashPassword(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

// CheckPasswordHash mocks the CheckPasswordHash method.
func (m *PasswordServiceMock) CheckPasswordHash(password, hash string) error {
	args := m.Called(password, hash)
	return args.Error(0)
}
