package apdata

func (c *Client) Login(username, password string) error {
	c.logger.Println("Logging in...")
	_, err := c.lowcli.Login(username, password)
	if err != nil {
		return err
	}
	c.logger.Println("Done!")
	return nil
}
