package main

import (
	service_api "gohtml/api"
	"gohtml/controller"
	"gohtml/infrastructure"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

func init() {
	infrastructure.NewSession()
}

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(logger.New())

	app.Static("/static", "./static")
	// app.Get("/", controller.Index)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("halo")
	})
	app.Get("/about", controller.About)
	app.Get("/layout", controller.Layout)
	app.Get("/login", controller.Login)

	api := app.Group("/api")
	api.Post("/login", service_api.Login)
	api.Get("/logout", service_api.Logout)

	log.Fatal(app.Listen(":3000"))
}
