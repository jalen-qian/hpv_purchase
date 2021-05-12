package main

import (
	"context"
	"github.com/gojektech/heimdall"
	"github.com/gojektech/heimdall/httpclient"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpClient struct {
	client *httpclient.Client
}

// NewHttpClient 创建 http client
func NewHttpClient(client heimdall.Doer) *HttpClient {
	backoff := heimdall.NewConstantBackoff(100*time.Millisecond, 300*time.Millisecond)
	retrier := heimdall.NewRetrier(backoff)
	timeout := 30 * time.Second
	retryCount := 3

	hcClient := httpclient.NewClient(
		httpclient.WithHTTPTimeout(timeout),
		httpclient.WithRetryCount(retryCount),
		httpclient.WithRetrier(retrier),
		httpclient.WithHTTPClient(client),
	)

	return &HttpClient{client: hcClient}
}

// HTTPProxyResponse 请求响应体
type HTTPProxyResponse struct {
	Status  int
	Body    []byte
	Headers http.Header
	Cookies []*http.Cookie
}

// doRequest 发起请求
func (c *HttpClient) doRequest(ctx context.Context,
	method,
	url string,
	header http.Header,
	cookies []http.Cookie,
	body io.Reader) (*HTTPProxyResponse, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header = header
	for _, ck := range cookies {
		req.AddCookie(&ck)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	rawBody := []byte("")
	if resp != nil {
		rawBody, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
	}
	return &HTTPProxyResponse{
		Status:  resp.StatusCode,
		Body:    rawBody,
		Headers: resp.Header,
		Cookies: resp.Cookies(),
	}, nil
}
