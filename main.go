package main

import (
	"api-validator/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	
    app := fiber.New()

	app.Use(logger.New(logger.Config{
        Format:     "[${ip}]:${port} ${status} - ${method} ${path}\n${body}",
	}))	

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Thank god it works 🙏")
    })

	app.Post("/", validator.ValidateAddDog, func(c *fiber.Ctx) error {
        body := new(validator.Dog)
        c.BodyParser(&body)
        return c.Status(fiber.StatusOK).JSON(body)
    })

    app.Listen(":2000")
}