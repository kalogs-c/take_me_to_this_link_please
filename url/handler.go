package url

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Post(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/html")

	original_url := c.FormValue("url")
	slug, err := HashURL(original_url)
	if err != nil {
		return c.SendString(err.Error())
	}

	return c.Render("components/new_url", fiber.Map{
		"Slug":    slug,
		"BaseURL": c.BaseURL(),
	})
}

func Redirect(c *fiber.Ctx) error {
	slug := c.Params("slug")
	url, err := VerifyURL(slug)
	fmt.Printf("url: %s\n slug: %s", url, slug)
	if err != nil {
		return c.SendString(err.Error())
	}

	return c.Redirect(url, 301)
}

func RenderIndex(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Take me to this link please",
	})
}

func Notify(c *fiber.Ctx) error {
	return c.Render("components/notification", nil)
}
