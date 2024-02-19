package main

import (
	"net/http"

	"github.com/go-resty/resty/v2"
)

func main() {
	restClient := resty.NewWithClient(new(http.Client))
	restClient.SetBaseURL("http://localhost:8080")
	resp, err := restClient.R().SetQueryParam("test", "").Get("/blocks")
	println(resp.Request.RawRequest.URL.RawQuery)
	if err != nil {
		panic(err)
	}
	println(resp.String())

}
