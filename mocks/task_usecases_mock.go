package mocks

import (
	"clean-architecture/domain"

	"github.com/stretchr/testify/mock"
)

// TaskUsecaseMock is a manually created mock for the TaskUsecase interface
type TaskUsecaseMock struct {
	mock.Mock
}

// GetTasks is a mock implementation of the TaskUsecase.GetTasks method
func (m *TaskUsecaseMock) GetTasks() ([]domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]domain.Task), args.Error(1)
}

// GetTaskByID is a mock implementation of the TaskUsecase.GetTaskByID method
func (m *TaskUsecaseMock) GetTaskByID(id string) (domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Task), args.Error(1)
}

// CreateTask is a mock implementation of the TaskUsecase.CreateTask method
func (m *TaskUsecaseMock) CreateTask(task domain.Task) (domain.Task, error) {
	args := m.Called(task)
	return args.Get(0).(domain.Task), args.Error(1)
}

// UpdateTask is a mock implementation of the TaskUsecase.UpdateTask method
func (m *TaskUsecaseMock) UpdateTask(id string, task domain.Task) (domain.Task, error) {
	args := m.Called(id, task)
	return args.Get(0).(domain.Task), args.Error(1)
}

// DeleteTask is a mock implementation of the TaskUsecase.DeleteTask method
func (m *TaskUsecaseMock) DeleteTask(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
