package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type LoginController interface {
	Login(c *fiber.Ctx) error
	Route(app *fiber.App)
}
