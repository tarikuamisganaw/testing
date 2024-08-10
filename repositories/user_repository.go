// Repositories/user_repository.go
package repositories

import (
	"context"
	"errors"
	"time"

	"clean-architecture/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Register(user domain.User) (domain.User, error)
	FindByUsername(username string) (domain.User, error)
	GetUsers() ([]domain.User, error)
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{
		collection: db.Collection("users"),
	}
}

func (ur *userRepository) Register(user domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := ur.collection.InsertOne(ctx, user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *userRepository) FindByUsername(username string) (domain.User, error) {
	var user domain.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := ur.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, errors.New("invalid username or password")
		}
		return user, err
	}

	return user, nil
}

func (ur *userRepository) GetUsers() ([]domain.User, error) {
	var users []domain.User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := ur.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user domain.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
