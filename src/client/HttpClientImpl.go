package client

import (
	"github.com/go-resty/resty/v2"
	"log"
)

func (client *HttpClientImpl) GET(queryParams map[string]string, url string) *resty.Response {
	response, err := client.client.R().SetQueryParams(queryParams).Get(url)
	if err != nil {
		log.Println("query error info=", err.Error())
	}
	return response
}
func (client *HttpClientImpl) POST(body interface{}, url string) *resty.Response {
	response, err := client.client.R().SetBody(body).Post(url)
	if err != nil {
		log.Println("create error info=", err.Error())
	}
	return response
}
func (client *HttpClientImpl) PUT(body interface{}, url string) *resty.Response {
	response, err := client.client.R().SetBody(body).Put(url)
	if err != nil {
		log.Println("update error info=", err.Error())
	}
	return response
}

func (client *HttpClientImpl) DELETE(params map[string]string, url string) *resty.Response {
	response, err := client.client.R().SetPathParams(params).Delete(url)
	if err != nil {
		log.Println("delete error info=", err.Error())
	}
	return response
}
