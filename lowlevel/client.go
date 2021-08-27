package lowlevel

import (
	"log"
	"net/http"
	"net/http/cookiejar"

	"github.com/HereMobilityDevelopers/mediary"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

type Client struct {
	logger    *log.Logger
	client    *resty.Client
	cookiejar http.CookieJar

	sessionID string
	ts        string
}

func NewClient(baseurl string, options ...ClientOption) (*Client, error) {
	builder := mediary.Init()
	builder.AddInterceptors(normalizeResponseJSON)

	optionsData := combineOptions(options)

	for _, i := range optionsData.interceptors {
		builder.AddInterceptors(i)
	}

	httpClient := builder.Build()

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create cookie jar")
	}

	return &Client{
		client: resty.NewWithClient(httpClient).
			SetHostURL(baseurl).
			SetCookieJar(jar),
		cookiejar: jar,
	}, nil
}
