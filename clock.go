package apdata

func (c *Client) Clock() (string, error) {
	c.logger.Println("Starting transaction...")
	resp, err := c.lowcli.GenericTransaction(true, map[string]string{
		"serverClass":  "ConBatidasTimerCalcEnabled",
		"sMessage":     "1895936000",
		"killAfterUse": "false",
	})
	if err != nil {
		c.logger.Fatalln(err)
		return "", err
	}
	c.logger.Println("Done!")

	c.logger.Println("Clocking...")
	resp, err = c.lowcli.ExecCalcLPCAndSaveTime(resp.Hwd)
	if err != nil {
		c.logger.Fatalln(err)
		return "", err
	}
	c.logger.Println("Done!")

	c.logger.Println(resp.Message)
	return resp.Message, nil
}
