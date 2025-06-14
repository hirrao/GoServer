package main

import (
	"GoServer/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	app := fiber.New()
	route.Routing(app)
	log.Error(app.Listen(":8121"))
}
