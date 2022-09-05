package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		BodyLimit:             4 * 1024 * 1024,
		Views:                 engine,
		ViewsLayout:           "",
		DisableStartupMessage: true,
		AppName:               "iamnotdoinganything",
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})

	log.Fatal(app.Listen(":8080"))
}
