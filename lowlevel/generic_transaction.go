package lowlevel

import "github.com/go-resty/resty/v2"

func (c *Client) GenericTransaction(post bool, data map[string]string) (*GenericResponse, error) {
	if data == nil {
		data = make(map[string]string)
	}

	data["tsc"] = c.ts
	data["sessionID"] = c.sessionID

	var response *resty.Response
	var err error
	if post {
		response, err = c.client.R().
			SetFormData(data).
			SetResult(GenericResponse{}).
			Post("/genericTransactionExecute")
	} else {
		response, err = c.client.R().
			SetQueryParams(data).
			SetResult(GenericResponse{}).
			Get("/genericTransactionExecute")
	}
	if err != nil {
		return nil, err
	}

	return c.handleResponse(response, "generic transaction")
}
