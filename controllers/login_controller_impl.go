package controllers

import (
	"github.com/Fadlihardiyanto/sppd-go/models"
	"github.com/Fadlihardiyanto/sppd-go/service"
	"github.com/gofiber/fiber/v2"
)

type LoginControllerImpl struct {
	UserService service.UserService
}

func NewLoginController(userService service.UserService) LoginController {
	return &LoginControllerImpl{
		UserService: userService,
	}
}

func (controller *LoginControllerImpl) Route(app *fiber.App) {
	app.Post("/api/login", controller.Login)
}

func (controller *LoginControllerImpl) Login(c *fiber.Ctx) error {
	var request models.LoginRequest
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	response, err := controller.UserService.Login(c.Context(), request)
	if err != nil {
		return err
	}

	return c.JSON(models.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})

}
