package controller

import (
	"fmt"
	"gohtml/infrastructure"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {

	sess, err := infrastructure.SessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	data := fiber.Map{
		"title": fmt.Sprintf("Hello %v", sess.Get("username")),
	}
	return c.Render("index", data)
}
