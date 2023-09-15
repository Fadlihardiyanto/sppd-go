package main

import (
	"github.com/Fadlihardiyanto/sppd-go/config"
	"github.com/Fadlihardiyanto/sppd-go/controllers"
	"github.com/Fadlihardiyanto/sppd-go/repository"
	"github.com/Fadlihardiyanto/sppd-go/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	database := config.GetConnection()

	// setup repository
	userRepository := repository.NewUserRepository(database)

	// setup service
	userService := service.NewUserService(database, userRepository, validator.New())

	// setup controller
	userController := controllers.NewUserController(userService)

	// setup fiber app
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	// setup route
	userController.Route(app)

	// run app
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}

}
