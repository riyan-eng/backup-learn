package controller

import "github.com/gofiber/fiber/v2"

func About(c *fiber.Ctx) error {
	data := fiber.Map{
		"title": "About",
	}
	return c.Render("about", data)
}
