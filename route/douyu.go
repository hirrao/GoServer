package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hirrao/go-server/common/http"
)

// RouteDouyu TODO 未完成不可使用
func RouteDouyu(app *fiber.App) {
	group := app.Group("/blive")
	group.Get("/:roomId", func(ctx *fiber.Ctx) error {
		res, err := getDouyuUrl(ctx.Params("roomId"))
		return http.SendResponse(res, err, ctx)
	})
}

func getDouyuUrl(roomId string) (interface{}, error) {
	return nil, nil
}
