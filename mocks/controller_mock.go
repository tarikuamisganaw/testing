package mocks

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

// TaskControllerMock is a mock implementation of the TaskController.
type TaskControllerMock struct {
	mock.Mock
}

func (m *TaskControllerMock) GetTasks(c *gin.Context) {
	m.Called(c)
}

func (m *TaskControllerMock) GetTaskByID(c *gin.Context) {
	m.Called(c)
}

func (m *TaskControllerMock) CreateTask(c *gin.Context) {
	m.Called(c)
}

func (m *TaskControllerMock) UpdateTask(c *gin.Context) {
	m.Called(c)
}

func (m *TaskControllerMock) DeleteTask(c *gin.Context) {
	m.Called(c)
}

// UserControllerMock is a mock implementation of the UserController.
type UserControllerMock struct {
	mock.Mock
}

func (m *UserControllerMock) Register(c *gin.Context) {
	m.Called(c)
}

func (m *UserControllerMock) Login(c *gin.Context) {
	m.Called(c)
}

func (m *UserControllerMock) GetUsers(c *gin.Context) {
	m.Called(c)
}
