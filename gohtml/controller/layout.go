package controller

import "github.com/gofiber/fiber/v2"

func Layout(c *fiber.Ctx) error {
	data := fiber.Map{
		"title": "Hello World",
	}
	return c.Render("index", data, "layouts/main")
}
