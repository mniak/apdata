package lowlevel

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

func (c *Client) handleResponse(resp *resty.Response, name string) (*GenericResponse, error) {
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("invalid status code %d", resp.StatusCode())
	}

	result, ok := resp.Result().(*GenericResponse)
	if !ok {
		return nil, fmt.Errorf("could not parse %s response", name)
	}
	if !result.Success {
		if result.Error == "" {
			return nil, fmt.Errorf("%s failed", name)
		}

		return nil, errors.Wrapf(errors.New(result.Error), "%s failed", name)
	}

	if newts := getCookie(resp, "ts"); newts != "" {
		c.ts = newts
	}
	return result, nil
}

func getCookie(r *resty.Response, name string) string {
	for _, cookie := range r.Cookies() {
		if cookie.Name == name {
			return cookie.Value
		}
	}
	return ""
}
