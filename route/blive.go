package route

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"strconv"
)

type body struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		PlayUrl struct {
			Durl []url `json:"durl"`
		} `json:"play_url"`
	} `json:"data"`
}

type url struct {
	Url string `json:"url"`
}

func RouteBLive(app *fiber.App) {
	group := app.Group("/blive")
	group.Get("/:roomId", func(ctx *fiber.Ctx) error {
		res, err := getUrl(ctx.Params("roomId"))
		if err != nil {
			return ctx.SendString(err.Error())
		}
		return ctx.JSON(res)
	})
}

func getUrl(rooId string) ([]url, error) {
	code, resp, err := fasthttp.Get(nil,
		fmt.Sprintf("https://api.live.bilibili.com/xlive/web-room/v1/index/getRoomPlayInfo?room_id=%s&play_url=1&mask=1&qn=20000&platform=web", rooId),
	)
	if err != nil {
		return nil, err
	}
	if code != 200 {
		return nil, errors.New(strconv.Itoa(code))

	}
	var b body
	err = json.Unmarshal(resp, &b)
	if err != nil {
		return nil, err
	}
	if b.Code != 0 {
		return nil, errors.New(b.Message)
	}
	return b.Data.PlayUrl.Durl, nil
}
