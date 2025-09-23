package data

import (
	"net/http"
)

type ClientMock struct {
	Response *http.Response
	Err      error
}

func (c *ClientMock) Do(req *http.Request) (*http.Response, error) {
    return c.Response, c.Err
}
