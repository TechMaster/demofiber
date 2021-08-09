package helper

import "github.com/gofiber/fiber/v2"

func ViewData(ctx *fiber.Ctx, key string, value interface{}) {
	if viewData, ok := ctx.Locals("ViewData").(fiber.Map); ok {
		viewData[key] = value
	} else {
		viewData = fiber.Map{
			key: value,
		}
		ctx.Locals("ViewData", viewData)
	}
}

func Render(ctx *fiber.Ctx, route string, viewData fiber.Map, layouts ...string) error {
	for k, v := range ctx.Locals("ViewData").(fiber.Map) {
		viewData[k] = v
	}
	return ctx.Render(route, viewData, layouts...)
}
