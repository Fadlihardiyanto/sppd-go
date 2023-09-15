package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	Create(c *fiber.Ctx) error
	Route(app *fiber.App)
	FindAll(c *fiber.Ctx) error
	FindByID(c *fiber.Ctx) error
}
