package main

import (
	"log"
	"tes/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", handler.Index)

	log.Fatal(app.Listen(":3000"))
}
