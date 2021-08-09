package controller

import (
	"demofiber/model"
	"fmt"

	"github.com/TechMaster/eris"
	"github.com/gofiber/fiber/v2"
)

func TestStruct(c *fiber.Ctx) error {
	base := model.Base{
		B:   10,
		Tag: "shop",
	}

	container := model.Container{
		Base: model.Base{
			B:   11,
			Tag: "sport",
		},
		Desc: "foo",
	}

	fmt.Println(base, container.Base.Tag)
	return nil
}

func PostJob(ctx *fiber.Ctx) error {
	var req model.Job
	if err := ctx.BodyParser(&req); err != nil {
		return eris.NewFrom(err).StatusCode(fiber.StatusBadRequest).EnableJSON()
	} else {
		return ctx.JSON(req)
	}
}
