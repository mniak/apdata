package lowlevel

func (c *Client) SetUserStates() (*GenericResponse, error) {
	formdata := map[string]string{
		"states":    "{\"btnDashboard\":0}",
		"tsc":       c.ts,
		"sessionID": c.sessionID,
	}

	response, err := c.client.R().
		SetFormData(formdata).
		SetResult(GenericResponse{}).
		Post("/setUserStates")
	if err != nil {
		return nil, err
	}

	return c.handleResponse(response, "get user states")
}
