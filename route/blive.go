package route

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/hirrao/go-server/common/http"
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
		res, err := getBliveUrl(ctx.Params("roomId"))
		if ctx.QueryInt("r", 0) == 1 {
			uri := res[0].Url
			return ctx.Redirect(uri, fiber.StatusFound)
		}
		return http.SendResponse(res, err, ctx)
	})
}

func getBliveUrl(roomId string) ([]url, error) {
	resp, err := http.Get(
		fmt.Sprintf("https://api.live.bilibili.com/xlive/web-room/v1/index/getRoomPlayInfo?room_id=%s&play_url=1&mask=1&qn=20000&platform=web", roomId),
	)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, errors.New(strconv.Itoa(resp.StatusCode()))

	}
	var b body
	err = json.Unmarshal(resp.Body(), &b)
	if err != nil {
		return nil, err
	}
	if b.Code != 0 {
		return nil, errors.New(b.Message)
	}
	if b.Data.PlayUrl.Durl == nil {
		return nil, errors.New("无法获取, 可能未开播")
	}
	return b.Data.PlayUrl.Durl, nil
}
