package repositories

import (
	"clean-architecture/domain"

	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository interface {
	GetTasks() ([]domain.Task, error)
	GetTaskByID(id string) (domain.Task, error)
	CreateTask(task domain.Task) (domain.Task, error)
	UpdateTask(id string, task domain.Task) (domain.Task, error)
	DeleteTask(id string) error
}

type taskRepository struct {
	collection *mongo.Collection
}

func NewTaskRepository(db *mongo.Database) TaskRepository {
	return &taskRepository{
		collection: db.Collection("tasks"),
	}
}

func (tr *taskRepository) GetTasks() ([]domain.Task, error) {
	var tasks []domain.Task
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := tr.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var task domain.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (tr *taskRepository) GetTaskByID(id string) (domain.Task, error) {
	var task domain.Task
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return task, errors.New("invalid task ID format")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = tr.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return task, errors.New("task not found")
		}
		return task, err
	}

	return task, nil
}

func (tr *taskRepository) CreateTask(task domain.Task) (domain.Task, error) {
	task.ID = primitive.NewObjectID()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := tr.collection.InsertOne(ctx, task)
	if err != nil {
		return task, err
	}

	return task, nil
}

func (tr *taskRepository) UpdateTask(id string, task domain.Task) (domain.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return task, errors.New("invalid task ID format")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := tr.collection.FindOneAndUpdate(ctx, bson.M{"_id": objID}, bson.M{"$set": task})
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return task, errors.New("task not found")
		}
		return task, result.Err()
	}

	task.ID = objID
	return task, nil
}

func (tr *taskRepository) DeleteTask(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid task ID format")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := tr.collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}

	return nil
}
