package main

import "github.com/gofiber/fiber/v2"

func main() {
	//Init Fiber
	app := fiber.New()

	//Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Ambasing!")
	})

	//Port Listening
	app.Listen(":3000")
}
