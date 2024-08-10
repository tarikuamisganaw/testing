package routers

import (
	"clean-architecture/delivery/controllers"
	"clean-architecture/infrastructure"
	"clean-architecture/usecases"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, taskUsecase usecases.TaskUsecase, userUsecase usecases.UserUsecase, jwtService infrastructure.JWTService) {
	taskController := controllers.NewTaskController(taskUsecase)
	userController := controllers.NewUserController(userUsecase)
	adminMiddleware := infrastructure.AdminMiddleware(jwtService)

	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)

	auth := r.Group("/api")
	auth.Use(infrastructure.AuthMiddleware(jwtService))
	{
		auth.GET("/tasks", taskController.GetTasks)
		auth.GET("/tasks/:id", taskController.GetTaskByID)
		auth.POST("/tasks", taskController.CreateTask)
		auth.PUT("/tasks/:id", taskController.UpdateTask)
		auth.DELETE("/tasks/:id", taskController.DeleteTask)
		auth.GET("/users", adminMiddleware, userController.GetUsers)
	}
}
