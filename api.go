package opensea

type Api struct {
	c *Client
}

func (a *Api) request(path string, post bool, params interface{}) IRequest {
	return &Request{
		Path:   path,
		Post:   post,
		Params: params,
	}
}

type IRequest interface {
	GetPath() string
	GetParams() interface{}
	IsPost() bool
}

type Request struct {
	Path   string
	Post   bool
	Params interface{}
}

func (r Request) GetPath() string {
	return r.Path
}

func (r Request) GetParams() interface{} {
	return r.Params
}

func (r Request) IsPost() bool {
	return r.Post
}

type IResponse interface{}
