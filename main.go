package main

import (
	"api-validator/validator"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

// use os package to get the env variable which is already set
func envVariable(key string) string {

    // set env variable using os package
    // os.Setenv(key, "gopher")
  
    // return the env variable using os package
    return os.Getenv(key)
}

func main() {
    godotenv.Load()
    value := envVariable("NAME")
    // log.Println(value)

    app := fiber.New()

	app.Use(logger.New(logger.Config{
        Format:     "[${ip}]:${port} ${status} - ${method} ${path}\n${body}",
	}))	

    app.Get("/", func(c *fiber.Ctx) error {
        
        return c.SendString("Thank god it works 🙏" + fmt.Sprintf("Read ENV : %s", value))
    })

	app.Post("/", validator.ValidateAddDog, func(c *fiber.Ctx) error {
        body := new(validator.Dog)
        c.BodyParser(&body)
        return c.Status(fiber.StatusOK).JSON(body)
    })

    app.Listen(":2000")
}