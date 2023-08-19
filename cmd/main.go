package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"github.com/kalogs-c/take_me_to_this_link_please/url"
)

func main() {
	engine := html.New("templates", ".html")
	config := fiber.Config{
		Views: engine,
	}
	app := fiber.New(config)

	app.Get("/", url.RenderIndex)
	app.Post("/url", url.Post)
	app.Get("/this-one/:slug", url.Redirect)
	app.Get("/notify", url.Notify)

	app.Listen(":8080")
}
