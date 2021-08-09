package controller

import "github.com/gofiber/fiber/v2"

func UseDefaultTemplate(c *fiber.Ctx) error {
	return c.Render("khoahoc/index", fiber.Map{
		"Title": "Khoá học làm giàu",
	})
}

func UseCustomTemplate(c *fiber.Ctx) error {
	return c.Render("khoahoc/index", fiber.Map{
		"Title": "Khoá học làm giàu",
	}, "layout/custom")
}

func UseNoTemplate(c *fiber.Ctx) error {
	return c.Render("khoahoc/index", fiber.Map{
		"Title": "Khoá học làm giàu",
	}, "_")
}
