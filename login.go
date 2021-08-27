package apdata

func (c *Client) Login(username, password string) error {
	_, err := c.lowcli.Login(username, password)
	return err
}
