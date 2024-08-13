package repositories

import (
	"clean-architecture/domain"
	"clean-architecture/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestTaskRepository_GetTasks(t *testing.T) {
	mockRepo := new(mocks.TaskRepositoryMock)
	taskRepo := mockRepo // Using the mock repository directly since it's a mock

	tasks := []domain.Task{
		{ID: primitive.NewObjectID(), Title: "Task 1"},
		{ID: primitive.NewObjectID(), Title: "Task 2"},
	}

	mockRepo.On("GetTasks").Return(tasks, nil)

	result, err := taskRepo.GetTasks()

	assert.NoError(t, err)
	assert.Equal(t, tasks, result)
	mockRepo.AssertExpectations(t)
}

func TestTaskRepository_GetTaskByID(t *testing.T) {
	mockRepo := new(mocks.TaskRepositoryMock)
	taskRepo := mockRepo // Using the mock repository directly since it's a mock

	taskID := primitive.NewObjectID().Hex()
	expectedTask := domain.Task{ID: primitive.NewObjectID(), Title: "Task 1"}

	mockRepo.On("GetTaskByID", taskID).Return(expectedTask, nil)

	result, err := taskRepo.GetTaskByID(taskID)

	assert.NoError(t, err)
	assert.Equal(t, expectedTask, result)
	mockRepo.AssertExpectations(t)
}

func TestTaskRepository_CreateTask(t *testing.T) {
	mockRepo := new(mocks.TaskRepositoryMock)
	taskRepo := mockRepo // Using the mock repository directly since it's a mock

	newTask := domain.Task{Title: "New Task"}
	mockRepo.On("CreateTask", mock.Anything).Return(newTask, nil)

	result, err := taskRepo.CreateTask(newTask)

	assert.NoError(t, err)
	assert.Equal(t, newTask, result)
	mockRepo.AssertExpectations(t)
}

func TestTaskRepository_UpdateTask(t *testing.T) {
	mockRepo := new(mocks.TaskRepositoryMock)
	taskRepo := mockRepo // Using the mock repository directly since it's a mock

	taskID := primitive.NewObjectID().Hex()
	updatedTask := domain.Task{Title: "Updated Task"}

	mockRepo.On("UpdateTask", taskID, mock.Anything).Return(updatedTask, nil)

	result, err := taskRepo.UpdateTask(taskID, updatedTask)

	assert.NoError(t, err)
	assert.Equal(t, updatedTask, result)
	mockRepo.AssertExpectations(t)
}

func TestTaskRepository_DeleteTask(t *testing.T) {
	mockRepo := new(mocks.TaskRepositoryMock)
	taskRepo := mockRepo // Using the mock repository directly since it's a mock

	taskID := primitive.NewObjectID().Hex()

	mockRepo.On("DeleteTask", taskID).Return(nil)

	err := taskRepo.DeleteTask(taskID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
