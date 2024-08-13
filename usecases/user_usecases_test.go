package usecases

import (
	"clean-architecture/domain"
	"clean-architecture/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserUsecase_Register(t *testing.T) {
	mockRepo := new(mocks.UserRepositoryMock)
	mockPasswordSvc := new(mocks.PasswordServiceMock)
	mockJwtSvc := new(mocks.JWTServiceMock)

	userUsecase := NewUserUsecase(mockRepo, mockPasswordSvc, mockJwtSvc)

	user := domain.User{Username: "testuser", Password: "plainpassword"}

	hashedPassword := "hashedpassword"
	mockPasswordSvc.On("HashPassword", user.Password).Return(hashedPassword, nil)

	// Set the expected user with hashed password
	expectedUser := user
	expectedUser.Password = hashedPassword

	mockRepo.On("Register", mock.Anything).Return(expectedUser, nil)

	result, err := userUsecase.Register(user)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, result)
	assert.Equal(t, hashedPassword, result.Password)
	mockRepo.AssertExpectations(t)
	mockPasswordSvc.AssertExpectations(t)
}

func TestUserUsecase_Login(t *testing.T) {
	mockRepo := new(mocks.UserRepositoryMock)
	mockPasswordSvc := new(mocks.PasswordServiceMock)
	mockJwtSvc := new(mocks.JWTServiceMock)

	userUsecase := NewUserUsecase(mockRepo, mockPasswordSvc, mockJwtSvc)

	username := "testuser"
	password := "plainpassword"
	hashedPassword := "hashedpassword"
	token := "generatedToken"

	user := domain.User{Username: username, Password: hashedPassword}

	mockRepo.On("FindByUsername", username).Return(user, nil)
	mockPasswordSvc.On("CheckPasswordHash", hashedPassword, password).Return(nil)
	mockJwtSvc.On("GenerateJWT", username, user.Role).Return(token, nil)

	result, err := userUsecase.Login(username, password)

	assert.NoError(t, err)
	assert.Equal(t, token, result)
	mockRepo.AssertExpectations(t)
	mockPasswordSvc.AssertExpectations(t)
	mockJwtSvc.AssertExpectations(t)
}

func TestUserUsecase_GetUsers(t *testing.T) {
	mockRepo := new(mocks.UserRepositoryMock)

	userUsecase := NewUserUsecase(mockRepo, nil, nil)

	users := []domain.User{
		{Username: "user1"},
		{Username: "user2"},
	}

	mockRepo.On("GetUsers").Return(users, nil)

	result, err := userUsecase.GetUsers()

	assert.NoError(t, err)
	assert.Equal(t, users, result)
	mockRepo.AssertExpectations(t)
}

func TestUserUsecase_Login_InvalidUsername(t *testing.T) {
	mockRepo := new(mocks.UserRepositoryMock)
	mockPasswordSvc := new(mocks.PasswordServiceMock)
	mockJwtSvc := new(mocks.JWTServiceMock)

	userUsecase := NewUserUsecase(mockRepo, mockPasswordSvc, mockJwtSvc)

	username := "invaliduser"
	password := "plainpassword"

	mockRepo.On("FindByUsername", username).Return(domain.User{}, errors.New("invalid username or password"))

	_, err := userUsecase.Login(username, password)

	assert.Error(t, err)
	assert.Equal(t, "invalid username or password", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestUserUsecase_Login_InvalidPassword(t *testing.T) {
	mockRepo := new(mocks.UserRepositoryMock)
	mockPasswordSvc := new(mocks.PasswordServiceMock)
	mockJwtSvc := new(mocks.JWTServiceMock)

	userUsecase := NewUserUsecase(mockRepo, mockPasswordSvc, mockJwtSvc)

	username := "testuser"
	password := "wrongpassword"
	hashedPassword := "hashedpassword"

	user := domain.User{Username: username, Password: hashedPassword}

	mockRepo.On("FindByUsername", username).Return(user, nil)
	mockPasswordSvc.On("CheckPasswordHash", hashedPassword, password).Return(errors.New("invalid username or password"))

	_, err := userUsecase.Login(username, password)

	assert.Error(t, err)
	assert.Equal(t, "invalid username or password", err.Error())
	mockRepo.AssertExpectations(t)
	mockPasswordSvc.AssertExpectations(t)
}
