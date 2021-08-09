package controller

import (
	"demofiber/helper"

	"github.com/gofiber/fiber/v2"
)

func Handler1(ctx *fiber.Ctx) error {
	helper.ViewData(ctx, "Data", "OX-13")
	helper.ViewData(ctx, "Counter", 100)
	return ctx.Next()
}

func Handler2(ctx *fiber.Ctx) error {
	ctx.Locals("ViewData").(fiber.Map)["Foo"] = true
	return ctx.Next()
}

func Handler3(ctx *fiber.Ctx) error {
	viewData := fiber.Map{
		"Title":   "This is test",
		"Content": "No fun"}

	return helper.Render(ctx, "democtx", viewData, "layout/custom")
}
