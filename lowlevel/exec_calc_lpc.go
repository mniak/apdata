package lowlevel

import "strconv"

func (c *Client) ExecCalcLPCAndSaveTime(hwd int64) (*GenericResponse, error) {
	formdata := map[string]string{
		"hwd":       strconv.Itoa(int(hwd)),
		"tsc":       c.ts,
		"sessionID": c.sessionID,
	}
	response, err := c.client.R().
		SetFormData(formdata).
		SetResult(GenericResponse{}).
		Post("/ExecCalcLPCAndSaveTime")
	if err != nil {
		return nil, err
	}

	return c.handleResponse(response, "clocking")
}
