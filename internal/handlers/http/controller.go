package http

import "github.com/gofiber/fiber/v2"

type Controller interface {
	Use(fiber.Router)
}
