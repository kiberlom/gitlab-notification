package servers

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/kiberlom/gitlab-notifications/docs" //lint:ignore golint
	"github.com/kiberlom/gitlab-notifications/internal/config"
	"github.com/kiberlom/gitlab-notifications/internal/handlers/http"
)

const APIPrefix = "api"

func NewHTTPServer(cfg *config.ConfigENV, controllers ...http.Controller) error {
	app := fiber.New()

	r := app.Group(fmt.Sprintf("/%s", APIPrefix))
	r.All("/_/docs/swagger/*", swagger.HandlerDefault)

	for _, c := range controllers {
		c.Use(r)
	}

	go func() {
		if err := app.Shutdown(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := app.Listen(fmt.Sprintf(":%d", cfg.ServerHTTPPort)); err != nil {
		return err
	}

	return nil
}
