package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Send([]byte("Welcome to Gobook"))
	})
	app.Use(cors.New())

	log.Fatal(app.Listen(":8080"))
}
