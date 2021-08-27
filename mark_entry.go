package apdata

func (c *Client) MarkEntry() error {
	c.logger.Println("Starting transaction")
	batidasResponse, err := c.lowcli.GenericTransaction(true, map[string]string{
		"serverClass":  "ConBatidasTimerCalcEnabled",
		"sMessage":     "1895936000",
		"killAfterUse": "false",
	})
	if err != nil {
		c.logger.Fatalln(err)
		return err
	}

	c.logger.Println("Marking entry")
	openResponse, err := c.lowcli.ExecCalcLPCAndSaveTime(batidasResponse.Hwd)
	if err != nil {
		c.logger.Fatalln(err)
		return err
	}

	c.logger.Println(openResponse.Message)
	return nil
}
