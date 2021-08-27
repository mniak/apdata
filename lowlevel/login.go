package lowlevel

import "fmt"

func (c *Client) Login(username, password string) (*GenericResponse, error) {
	formdata := map[string]string{
		"username": username,
		"password": password,
	}
	response, err := c.client.R().
		SetFormData(formdata).
		SetResult(GenericResponse{}).
		Post("/login")
	if err != nil {
		return nil, err
	}

	if !response.IsSuccess() {
		return nil, fmt.Errorf("invalid status code %d", response.StatusCode())
	}

	return c.handleResponse(response, "login")
}
