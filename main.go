package main

import (
	"go-react-example-api/controller"
	"go-react-example-api/db"
	"go-react-example-api/repository"
	"go-react-example-api/router"
	"go-react-example-api/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUseCase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
