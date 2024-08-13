package usecases

import (
	"clean-architecture/domain"
	"clean-architecture/repositories"
)

type TaskUsecase interface {
	GetTasks() ([]domain.Task, error)
	GetTaskByID(id string) (domain.Task, error)
	CreateTask(task domain.Task) (domain.Task, error)
	UpdateTask(id string, task domain.Task) (domain.Task, error)
	DeleteTask(id string) error
}

type taskUsecase struct {
	taskRepo repositories.TaskRepository
}

func NewTaskUsecase(tr repositories.TaskRepository) TaskUsecase {
	return &taskUsecase{taskRepo: tr}
}

// GetTasks retrieves all tasks
func (u *taskUsecase) GetTasks() ([]domain.Task, error) {
	tasks, err := u.taskRepo.GetTasks()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// GetTaskByID retrieves a task by its ID
func (u *taskUsecase) GetTaskByID(id string) (domain.Task, error) {
	task, err := u.taskRepo.GetTaskByID(id)
	if err != nil {
		return domain.Task{}, err
	}
	return task, nil
}

// CreateTask creates a new task
func (u *taskUsecase) CreateTask(task domain.Task) (domain.Task, error) {
	createdTask, err := u.taskRepo.CreateTask(task)
	if err != nil {
		return domain.Task{}, err
	}
	return createdTask, nil
}

// UpdateTask updates an existing task
func (u *taskUsecase) UpdateTask(id string, task domain.Task) (domain.Task, error) {
	updatedTask, err := u.taskRepo.UpdateTask(id, task)
	if err != nil {
		return domain.Task{}, err
	}
	return updatedTask, nil
}

// DeleteTask deletes a task by its ID
func (u *taskUsecase) DeleteTask(id string) error {
	if err := u.taskRepo.DeleteTask(id); err != nil {
		return err
	}
	return nil
}
