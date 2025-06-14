package route

import "github.com/gofiber/fiber/v2"

func Routing(app *fiber.App) {
	RouteBLive(app)
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, World")
	})
}
