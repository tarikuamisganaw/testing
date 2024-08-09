package main

import (
	"clean-architecture/delivery/routers"
	"clean-architecture/infrastructure"
	"clean-architecture/repositories"
	"clean-architecture/usecases"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("taskdb")

	taskRepo := repositories.NewTaskRepository(db)
	userRepo := repositories.NewUserRepository(db)

	// Create instances of PasswordService and JWTService
	passwordService := infrastructure.NewPasswordService() // assuming NewPasswordService exists
	secretKey := "your-secret-key"
	jwtService := infrastructure.NewJWTService(secretKey)

	// Pass PasswordService and JWTService to NewUserUsecase
	taskUsecase := usecases.NewTaskUsecase(taskRepo)
	userUsecase := usecases.NewUserUsecase(userRepo, passwordService, jwtService)

	r := gin.Default()
	routers.InitRoutes(r, taskUsecase, userUsecase, jwtService)

	r.Run(":8080")
}
