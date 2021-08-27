package apdata

import (
	"log"

	"github.com/mniak/apdata/lowlevel"
)

type Client struct {
	lowcli *lowlevel.Client
	logger *log.Logger
}

func NewClient(baseurl string, options ...ClientOption) (*Client, error) {
	optionsData := combineOptions(options)

	lowLevelClient, err := lowlevel.NewClient(baseurl,
		lowlevel.WithDebug(optionsData.debug),
		lowlevel.WithInterceptors(optionsData.interceptors...),
	)
	if err != nil {
		return nil, err
	}

	return &Client{
		lowcli: lowLevelClient,
		logger: optionsData.logger,
	}, nil
}
