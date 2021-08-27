package lowlevel

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/HereMobilityDevelopers/mediary"
	"github.com/mniak/semijson"
	"github.com/pkg/errors"
)

func normalizeResponseJSON(req *http.Request, handler mediary.Handler) (*http.Response, error) {
	resp, err := handler(req)
	if err != nil {
		return nil, err
	}

	originalBody := resp.Body
	defer originalBody.Close()

	jvalue, err := semijson.Parse(originalBody)
	if err != nil {
		return resp, errors.Wrap(err, "failed to normalize JSON")
	}

	resp.Header.Set("Content-Type", "application/json; charset=utf-8")

	json := jvalue.JSON()
	newbody := ioutil.NopCloser(bytes.NewBufferString(json))
	resp.Body = newbody

	return resp, err
}
