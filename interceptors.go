package apdata

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/HereMobilityDevelopers/mediary"
)

func DumpInterceptor(req *http.Request, handler mediary.Handler) (*http.Response, error) {
	if bytes, err := httputil.DumpRequestOut(req, true); err == nil {
		fmt.Printf("%s", bytes)
	}
	resp, err := handler(req)
	if bytes, err := httputil.DumpResponse(resp, true); err == nil {
		fmt.Printf("%s", bytes)
	}
	return resp, err
}
