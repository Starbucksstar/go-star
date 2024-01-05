package client

import "github.com/go-resty/resty/v2"

type HttpClient interface {
	GET(queryParams map[string]string, url string) *resty.Response
	POST(body interface{}, url string) *resty.Response
	PUT(body interface{}, url string) *resty.Response
	DELETE(params map[string]string, url string) *resty.Response
}

type HttpClientImpl struct {
	client *resty.Client
}

func NewHttpClient() HttpClient {
	return &HttpClientImpl{
		client: resty.New(),
	}
}
