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
		Timeout: time.Second * 10,
	}
)

func httpReqWithRetry(urlS string) (_ *http.Response, err error) {
	for i := 0; i < MaxHttpRetries; i++ {
		req, err := http.NewRequest(http.MethodGet, urlS, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Add("user-agent", "github.com/nektro/go-util/mbpp")
		req.Header.Add("connection", "close")
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
