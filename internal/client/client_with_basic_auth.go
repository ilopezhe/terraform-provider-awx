package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type clientWithBasicAuth struct {
	client *http.Client

	username, password, hostname string
	version                      string
}

var _ Client = &clientWithBasicAuth{}

func NewClientWithBasicAuth(username, password, hostname string, version string, insecureSkipVerify bool, httpClient *http.Client) Client {
	return &clientWithBasicAuth{
		client:   defaultClient(httpClient, insecureSkipVerify),
		hostname: hostname,
		username: username,
		password: password,
		version:  version,
	}
}

func (c *clientWithBasicAuth) NewRequest(ctx context.Context, method string, endpoint string, body io.Reader) (req *http.Request, err error) {
	endpoint = strings.TrimPrefix(endpoint, "/")
	req, err = http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s/%s", c.hostname, endpoint), body)
	if err == nil {
		req.SetBasicAuth(c.username, c.password)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", fmt.Sprintf("terraform-provider-awx/%s", c.version))
	}
	return req, err
}

func (c *clientWithBasicAuth) Do(ctx context.Context, req *http.Request) (data map[string]any, err error) {
	return doRequest(c.client, ctx, req)
}

// IsResourceNotFound checks if the given error indicates a 404 Not Found response.
func (c *clientWithBasicAuth) IsResourceNotFound(err error) bool {
	if err == nil {
		return false
	}
	// Check if the error message contains "status code: 404"
	return strings.Contains(err.Error(), "status code: 404") || strings.Contains(err.Error(), "status code: 400")
}
