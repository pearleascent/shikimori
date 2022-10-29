package shikimori

import (
	"errors"
	"io"
	"net/http"
	"strings"
)

func (c *Client) makeRequest(url string) ([]byte, error) {
	req, reqErr := http.NewRequest("GET", url, nil)
	if reqErr != nil {
		return []byte{}, reqErr
	}

	res, resErr := c.client.Do(req)
	if resErr != nil {
		return []byte{}, resErr
	}
	defer res.Body.Close()

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return []byte{}, readErr
	}

	if strings.Contains(string(body), "Retry later") {
		return []byte{}, errors.New("rate limiter triggered")
	}

	return body, nil
}
