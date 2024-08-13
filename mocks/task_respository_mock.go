// mocks/task_repository_mock.go
package mocks

import (
	"clean-architecture/domain"

	"github.com/stretchr/testify/mock"
)

// TaskRepositoryMock is a manually created mock for the TaskRepository interface
type TaskRepositoryMock struct {
	mock.Mock
}

// GetTasks is a mock implementation of the TaskRepository.GetTasks method
func (m *TaskRepositoryMock) GetTasks() ([]domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]domain.Task), args.Error(1)
}

// GetTaskByID is a mock implementation of the TaskRepository.GetTaskByID method
func (m *TaskRepositoryMock) GetTaskByID(id string) (domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Task), args.Error(1)
}

// CreateTask is a mock implementation of the TaskRepository.CreateTask method
func (m *TaskRepositoryMock) CreateTask(task domain.Task) (domain.Task, error) {
	args := m.Called(task)
	return args.Get(0).(domain.Task), args.Error(1)
}

// UpdateTask is a mock implementation of the TaskRepository.UpdateTask method
func (m *TaskRepositoryMock) UpdateTask(id string, task domain.Task) (domain.Task, error) {
	args := m.Called(id, task)
	return args.Get(0).(domain.Task), args.Error(1)
}

// DeleteTask is a mock implementation of the TaskRepository.DeleteTask method
func (m *TaskRepositoryMock) DeleteTask(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
