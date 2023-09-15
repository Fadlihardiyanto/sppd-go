package controllers

import (
	"github.com/Fadlihardiyanto/sppd-go/models"
	"github.com/Fadlihardiyanto/sppd-go/service"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Route(app *fiber.App) {
	app.Post("/api/user", controller.Create)
	app.Get("/api/user", controller.FindAll)
	app.Get("/api/user/:id", controller.FindByID)
}

func (controller *UserControllerImpl) Create(c *fiber.Ctx) error {
	var request models.CreateUserRequest
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	// generate uuid
	request.IDUser = uuid.New().String()

	//bcrypt password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// set password
	request.Password = string(hashedPassword)

	response, err := controller.UserService.Create(c.Context(), request)
	if err != nil {
		return err
	}

	return c.JSON(models.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *UserControllerImpl) FindAll(c *fiber.Ctx) error {
	response, err := controller.UserService.FindAll(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(models.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *UserControllerImpl) FindByID(c *fiber.Ctx) error {
	id := c.Params("id")
	response, err := controller.UserService.FindByID(c.Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(models.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}
