// usecases/task_usecase_test.go
package usecases

import (
	"clean-architecture/domain"
	"clean-architecture/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestTaskUsecase_GetTasks(t *testing.T) {
	mockRepo := new(mocks.TaskRepositoryMock)
	taskUsecase := NewTaskUsecase(mockRepo)

	tasks := []domain.Task{
		{ID: primitive.NewObjectID(), Title: "Task 1"},
		{ID: primitive.NewObjectID(), Title: "Task 2"},
	}

	mockRepo.On("GetTasks").Return(tasks, nil)

	result, err := taskUsecase.GetTasks()

	assert.NoError(t, err)
	assert.Equal(t, tasks, result)
	mockRepo.AssertExpectations(t)
}

func TestTaskUsecase_GetTaskByID(t *testing.T) {
	mockRepo := new(mocks.TaskRepositoryMock)
	taskUsecase := NewTaskUsecase(mockRepo)

	taskID := primitive.NewObjectID().Hex()
	expectedTask := domain.Task{ID: primitive.NewObjectID(), Title: "Task 1"}

	mockRepo.On("GetTaskByID", taskID).Return(expectedTask, nil)

	result, err := taskUsecase.GetTaskByID(taskID)

	assert.NoError(t, err)
	assert.Equal(t, expectedTask, result)
	mockRepo.AssertExpectations(t)
}

func TestTaskUsecase_CreateTask(t *testing.T) {
	mockRepo := new(mocks.TaskRepositoryMock)
	taskUsecase := NewTaskUsecase(mockRepo)

	newTask := domain.Task{Title: "New Task"}
	mockRepo.On("CreateTask", mock.Anything).Return(newTask, nil)

	result, err := taskUsecase.CreateTask(newTask)

	assert.NoError(t, err)
	assert.Equal(t, newTask, result)
	mockRepo.AssertExpectations(t)
}

func TestTaskUsecase_UpdateTask(t *testing.T) {
	mockRepo := new(mocks.TaskRepositoryMock)
	taskUsecase := NewTaskUsecase(mockRepo)

	taskID := primitive.NewObjectID().Hex()
	updatedTask := domain.Task{Title: "Updated Task"}

	mockRepo.On("UpdateTask", taskID, mock.Anything).Return(updatedTask, nil)

	result, err := taskUsecase.UpdateTask(taskID, updatedTask)

	assert.NoError(t, err)
	assert.Equal(t, updatedTask, result)
	mockRepo.AssertExpectations(t)
}

func TestTaskUsecase_DeleteTask(t *testing.T) {
	mockRepo := new(mocks.TaskRepositoryMock)
	taskUsecase := NewTaskUsecase(mockRepo)

	taskID := primitive.NewObjectID().Hex()

	mockRepo.On("DeleteTask", taskID).Return(nil)

	err := taskUsecase.DeleteTask(taskID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
