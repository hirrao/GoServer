package http

import "github.com/gofiber/fiber/v2"

func SendResponse(data interface{}, err error, ctx *fiber.Ctx) error {
	if err != nil {
		return ctx.SendString(err.Error())
	}
	return ctx.JSON(data)
}
