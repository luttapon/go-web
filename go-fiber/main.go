package main

import (
	"github.com/gofiber/fiber/v3"
)

// go run main.go
func main() {
	 app := fiber.New()

	 app.Get("/", func(c fiber.Ctx) error {
		 return c.SendString("Hello, World!")
	 })
	 app.Get("/", func(c fiber.Ctx) error {
		 return c.SendString("Hello, World!")
	 })
	 // start server
	 if err := app.Listen(":3000"); err != nil {
		 panic(err)
	 }
}