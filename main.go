package main

import (
	"fmt"

	"github.com/gofiber/fiber"
)

func main() {
	fmt.Println("main.go")

	app := fiber.New()

	//GET Requests
	app.Get("/api", func(c *fiber.Ctx) {
		c.Send("Hello World!")
	})
	//GET Request using parameters
	app.Get("/api/:id", func(c *fiber.Ctx) {
		id := c.Params("id")
		c.Send("Query parameter is " + id)
	})

	app.Listen(3000)
}
