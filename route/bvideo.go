package route

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"strconv"
)

type Response struct {
	code    int
	message string
	ttl     int
	data    interface{}
}

type Video struct {
	video string
	audio string
}

// 未完成, 不可使用
// 暂时放弃后续计划, 因为未登录无法获取720p以上视频
func routeBvideo(app *fiber.App) {
	type Video struct {
	}
	group := app.Group("/bvideo")
	group.Get("/:bvid", func(ctx *fiber.Ctx) error {
		_, err := getBiliVideo(ctx.Params("bvid"))
		if err != nil {
			return ctx.Status(500).SendString(err.Error())
		}
		return ctx.Status(500).SendString("TODO")
	})
}

func getBiliVideo(bvid string) (interface{}, error) {
	res, err := getCvid(bvid)
	if err != nil {
		return res, err
	}
	avid := res.(string)
	_, err = getDash(avid, bvid)
	return nil, errors.New("TODO")
}

func getDash(bvid string, cvid string) (interface{}, error) {
	code, resp, err := fasthttp.Get(nil,
		fmt.Sprintf("https://api.bilibili.com/x/player/playurl?qn=120&otype=json&fourk=1&fnver=0&fnval=4048&bvid=%s&cid%s", bvid, cvid),
	)
	if err != nil {
		return nil, err
	}
	if code != 200 {
		return nil, errors.New(strconv.Itoa(code))
	}
	var result Response
	var _ map[string]interface{}
	_ = json.Unmarshal(resp, &result)
	if result.data == nil {
		return nil, errors.New(result.message)
	}
	return nil, errors.New("TODO")
}

func getCvid(bvid string) (interface{}, error) {
	code, resp, err := fasthttp.Get(nil,
		fmt.Sprintf("https://api.bilibili.com/x/web-interface/view?bvid=%s", bvid))
	if err != nil {
		return nil, err
	}
	if code != 200 {
		return nil, errors.New(strconv.Itoa(code))
	}
	var result Response
	var data map[string]interface{}
	_ = json.Unmarshal(resp, &result)
	if result.data == nil {
		return nil, errors.New(result.message)
	}
	return strconv.Itoa(int(data["cid"].(float64))), nil
}
