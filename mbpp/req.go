package mbpp

import (
	"net/http"
	"strings"
	"time"
)

var (
	MaxHttpRetries = 3
)

var (
	netClient = &http.Client{
		Transport: &http.Transport{
			ResponseHeaderTimeout: time.Second * 10,
		},
	}
)

func httpReqWithRetry(urlS string, headers map[string]string) (_ *http.Response, err error) {
	for i := 0; i < MaxHttpRetries; i++ {
		req, err := http.NewRequest(http.MethodGet, urlS, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("user-agent", "github.com/nektro/go-util/mbpp")
		req.Header.Set("connection", "close")
		req.Header.Set("host", req.URL.Host)
		if headers != nil {
			for k, v := range headers {
				req.Header.Set(k, v)
			}
		}
		res, err := netClient.Do(req)
		if err != nil {
			if strings.Contains(err.Error(), "Client.Timeout exceeded") {
				continue
			}
		}
		return res, err
	}
	return nil, err
}
