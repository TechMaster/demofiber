package main

import (
	"demofiber/cookie_session"
	"demofiber/router"
	"demofiber/template"
	"fmt"

	"github.com/gofiber/fiber/v2"
) //import thư viện fiber version 2

func main() {
	app := fiber.New(
		fiber.Config{
			Views:        template.Init(),
			ErrorHandler: CustomErrorHandler, //Đăng ký hàm xử lý lỗi ở đây
		},
	)
	redisDB := cookie_session.InitSession()
	defer redisDB.Close()

	router.RegisterRoutes(app)
	//Để hàm này dưới cùng để bắt lỗi 404 Not Found
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Đường dẫn không tìm thấy")
	})

	if err := app.Listen(":3000"); err != nil {
		fmt.Println(err.Error())
	}
}
