package route

/*

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

func routeBvideo(app fiber.App) {
	type Video struct {
	}
	group := app.Group("/bvideo")
	group.Get("/:bvid", func(ctx *fiber.Ctx) error {
		resp, err := getVideo(ctx.Params("bvid"))
		if err != nil {
			return ctx.Status(500).SendString(err.Error())
		}
	})
}

func getVideo(bvid string) (interface{}, error) {
	res, err := getCvid(bvid)
	if err != nil {
		return res, err
	}
	avid := res.(string)
	dash, err := getDash(avid, bvid)
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
	var data map[string]interface{}
	_ = json.Unmarshal(resp, &result)
	if result.data == nil {
		return nil, errors.New(result.message)
	}

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
*/
