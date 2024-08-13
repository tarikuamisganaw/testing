// repositories/user_repository_test.go
package repositories

import (
	"clean-architecture/domain"
	"clean-architecture/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserRepository_Register(t *testing.T) {
	mockRepo := new(mocks.UserRepositoryMock)

	user := domain.User{
		ID:       primitive.NewObjectID(),
		Username: "testuser",
		Password: "password123",
	}

	mockRepo.On("Register", mock.Anything).Return(user, nil)

	result, err := mockRepo.Register(user)

	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockRepo.AssertExpectations(t)
}

func TestUserRepository_FindByUsername(t *testing.T) {
	mockRepo := new(mocks.UserRepositoryMock)

	username := "testuser"
	expectedUser := domain.User{
		ID:       primitive.NewObjectID(),
		Username: "testuser",
		Password: "password123",
	}

	mockRepo.On("FindByUsername", username).Return(expectedUser, nil)

	result, err := mockRepo.FindByUsername(username)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, result)
	mockRepo.AssertExpectations(t)
}

func TestUserRepository_GetUsers(t *testing.T) {
	mockRepo := new(mocks.UserRepositoryMock)

	users := []domain.User{
		{ID: primitive.NewObjectID(), Username: "user1", Password: "pass1"},
		{ID: primitive.NewObjectID(), Username: "user2", Password: "pass2"},
	}

	mockRepo.On("GetUsers").Return(users, nil)

	result, err := mockRepo.GetUsers()

	assert.NoError(t, err)
	assert.Equal(t, users, result)
	mockRepo.AssertExpectations(t)
}
