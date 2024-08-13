// Domain/domain.go
package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TaskUsecase defines the behavior for interacting with tasks
type TaskUsecase interface {
	GetTasks() ([]Task, error)
	GetTaskByID(string) (Task, error)
	CreateTask(Task) (Task, error)
	UpdateTask(string, Task) (Task, error)
	DeleteTask(string) error
}

// UserUsecase defines the behavior for interacting with users
type UserUsecase interface {
	Register(User) (User, error)
	Login(string, string) (string, error)
	GetUsers() ([]User, error)
}

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `json:"title" binding:"required"`
	Description string             `json:"description"`
	DueDate     time.Time          `json:"due_date"`
	Status      string             `json:"status"`
}

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `json:"username" binding:"required"`
	Password string             `json:"password" binding:"required"`
	Role     string             `json:"role"`
}
