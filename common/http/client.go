package http

import (
	"github.com/valyala/fasthttp"
	"time"
)

var Client = &fasthttp.Client{
	DialDualStack: false,
	ReadTimeout:   10 * time.Second,
}

func Get(url string) (*fasthttp.Response, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	req.SetRequestURI(url)
	req.Header.SetMethod("GET")

	err := Client.Do(req, resp)
	fasthttp.ReleaseRequest(req)
	if err != nil {
		fasthttp.ReleaseResponse(resp)
		return nil, err
	}
	return resp, nil
}

func Post(url string, body []byte) (*fasthttp.Response, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	req.SetRequestURI(url)
	req.Header.SetMethod("POST")
	req.SetBody(body)

	err := Client.Do(req, resp)
	fasthttp.ReleaseRequest(req)
	if err != nil {
		fasthttp.ReleaseResponse(resp)
		return nil, err
	}
	return resp, nil
}
