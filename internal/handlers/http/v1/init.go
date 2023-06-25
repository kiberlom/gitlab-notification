package v1

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kiberlom/gitlab-notifications/internal/handlers/http"
)

const version = "v1"

type HandlerSetup struct {
}

type resource struct {
}

func NewHandler(setup *HandlerSetup) http.Controller {
	return &resource{}
}

func (r *resource) Use(router fiber.Router) {
	version := router.Group(fmt.Sprintf("%s/", version))
	version.Get("/time", func(c *fiber.Ctx) error {
		return c.SendString(time.Now().String())
	})
}
