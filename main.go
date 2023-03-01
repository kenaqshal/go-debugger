package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Docker! <3")
	})

	app.Get("/test", func(c *fiber.Ctx) error {
		message := "Hi Everyone!"
		datetime := time.Now().Format("2006-01-02 15:04:05")
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": message,
			"date":    datetime,
		})
	})

	app.Listen(":3000")
}
