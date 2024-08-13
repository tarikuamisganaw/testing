package controllers

import (
	"clean-architecture/domain"
	"clean-architecture/mocks"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserController_Register(t *testing.T) {
	mockUserUsecase := new(mocks.UserUsecaseMock)
	userController := NewUserController(mockUserUsecase)

	objectID := primitive.NewObjectID() // Generate a new ObjectID
	user := domain.User{ID: objectID, Username: "testuser", Password: "password"}
	registeredUser := user

	mockUserUsecase.On("Register", user).Return(registeredUser, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/register", strings.NewReader(`{"username":"testuser","password":"password"}`))
	c.Request.Header.Set("Content-Type", "application/json")

	userController.Register(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	var responseUser domain.User
	json.Unmarshal(w.Body.Bytes(), &responseUser)
	assert.Equal(t, registeredUser, responseUser)
	mockUserUsecase.AssertExpectations(t)
}

func TestUserController_Login(t *testing.T) {
	mockUserUsecase := new(mocks.UserUsecaseMock)
	userController := NewUserController(mockUserUsecase)

	user := domain.User{Username: "testuser", Password: "password"}
	token := "mockToken"

	mockUserUsecase.On("Login", user.Username, user.Password).Return(token, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/login", strings.NewReader(`{"username":"testuser","password":"password"}`))
	c.Request.Header.Set("Content-Type", "application/json")

	userController.Login(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, token, response["token"])
	mockUserUsecase.AssertExpectations(t)
}

func TestUserController_GetUsers(t *testing.T) {
	mockUserUsecase := new(mocks.UserUsecaseMock)
	userController := NewUserController(mockUserUsecase)

	users := []domain.User{
		{ID: primitive.NewObjectID(), Username: "user1"},
		{ID: primitive.NewObjectID(), Username: "user2"},
	}

	mockUserUsecase.On("GetUsers").Return(users, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/users", nil)

	userController.GetUsers(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var responseUsers []domain.User
	json.Unmarshal(w.Body.Bytes(), &responseUsers)
	assert.Equal(t, users, responseUsers)
	mockUserUsecase.AssertExpectations(t)
}

func TestTaskController_GetTasks(t *testing.T) {
	mockTaskUsecase := new(mocks.TaskUsecaseMock)
	taskController := NewTaskController(mockTaskUsecase)

	tasks := []domain.Task{
		{ID: primitive.NewObjectID(), Title: "Task 1"},
		{ID: primitive.NewObjectID(), Title: "Task 2"},
	}

	mockTaskUsecase.On("GetTasks").Return(tasks, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/tasks", nil)

	taskController.GetTasks(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var responseTasks []domain.Task
	json.Unmarshal(w.Body.Bytes(), &responseTasks)
	assert.Equal(t, tasks, responseTasks)
	mockTaskUsecase.AssertExpectations(t)
}

func TestTaskController_CreateTask(t *testing.T) {
	mockTaskUsecase := new(mocks.TaskUsecaseMock)
	taskController := NewTaskController(mockTaskUsecase)

	objectID := primitive.NewObjectID() // Generate a new ObjectID
	task := domain.Task{ID: objectID, Title: "New Task"}
	createdTask := task

	mockTaskUsecase.On("CreateTask", task).Return(createdTask, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/tasks", strings.NewReader(`{"title":"New Task"}`))
	c.Request.Header.Set("Content-Type", "application/json")

	taskController.CreateTask(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	var responseTask domain.Task
	json.Unmarshal(w.Body.Bytes(), &responseTask)
	assert.Equal(t, createdTask, responseTask)
	mockTaskUsecase.AssertExpectations(t)
}
