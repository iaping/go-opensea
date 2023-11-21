package opensea

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
)

const (
	host    = "https://api.opensea.io"
	timeout = 10 * time.Second
)

type Client struct {
	host string
	key  string
	c    *fasthttp.Client
}

func New(key string, c *fasthttp.Client) *Client {
	if c == nil {
		c = DefaultFasthttpClient()
	}

	return &Client{
		host: host,
		key:  key,
		c:    c,
	}
}

func DefaultFasthttpClient() *fasthttp.Client {
	return &fasthttp.Client{
		Name:                "iaping/go-opensea",
		MaxConnsPerHost:     16,
		MaxIdleConnDuration: 20 * time.Second,
		ReadTimeout:         10 * time.Second,
		WriteTimeout:        10 * time.Second,
	}
}

func (c *Client) Do(req IRequest, resp IResponse) error {
	url, body, err := c.format(req)
	if err != nil {
		return err
	}

	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(request)
		fasthttp.ReleaseResponse(response)
	}()

	request.SetRequestURI(url)
	request.Header.Set("x-api-key", c.key)
	request.Header.Set("accept", "application/json")
	if body != nil {
		request.Header.Set("content-type", "application/json")
		request.SetBody(body)
	}

	if err := c.c.Do(request, response); err != nil {
		return errors.Wrap(err, url)
	}
	if response.StatusCode() != fasthttp.StatusOK {
		return fmt.Errorf("url: %s, code: %d", url, response.StatusCode())
	}

	return json.Unmarshal(response.Body(), &resp)
}

func (c *Client) format(req IRequest) (string, []byte, error) {
	url := c.host + req.GetPath()
	if req.IsPost() {
		body, err := json.Marshal(req.GetParams())
		if err != nil {
			return "", nil, err
		}
		return url, body, nil
	}

	v, err := query.Values(req.GetParams())
	if err != nil {
		return "", nil, err
	}
	return fmt.Sprintf("%s?%s", url, v.Encode()), nil, nil
}

func (c *Client) GetListings() *GetListings {
	return NewGetListings(c)
}
