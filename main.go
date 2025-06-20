package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/hirrao/go-server/route"
)

func main() {
	app := fiber.New()
	route.Routing(app)
	log.Error(app.Listen(":8121"))
}
