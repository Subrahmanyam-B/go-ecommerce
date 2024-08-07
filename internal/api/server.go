package api

import (
	"go-ecommerce-app/config"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func StartSever(config config.AppConfig) {

	app := fiber.New()

	app.Get("/health", HealthCheck)

	app.Listen(config.ServerPort)

}

func HealthCheck(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "I'm breathing!",
	})
}
