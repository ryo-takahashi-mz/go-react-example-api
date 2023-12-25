package main

import (
	"go-react-example-api/controller"
	"go-react-example-api/db"
	"go-react-example-api/repository"
	"go-react-example-api/router"
	"go-react-example-api/usecase"
	"go-react-example-api/validator"
)

func main() {
	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	userUseCase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUseCase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	userController := controller.NewUserController(userUseCase)
	taskController := controller.NewTaskController(taskUseCase)
	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
