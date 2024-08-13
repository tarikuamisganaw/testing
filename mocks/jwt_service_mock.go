package mocks

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/mock"
)

// JWTServiceMock is a mock implementation of the JWTService interface.
type JWTServiceMock struct {
	mock.Mock
}

// GenerateJWT mocks the GenerateJWT method.
func (m *JWTServiceMock) GenerateJWT(username, role string) (string, error) {
	args := m.Called(username, role)
	return args.String(0), args.Error(1)
}

// ValidateJWT mocks the ValidateJWT method.
func (m *JWTServiceMock) ValidateJWT(tokenString string) (*jwt.Token, error) {
	args := m.Called(tokenString)
	return args.Get(0).(*jwt.Token), args.Error(1)
}
