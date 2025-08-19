package scheduler

import (
	"fmt"
	"io"
	"net/http"
)

func DownloadFile(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("get err %s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status code: %d %s", resp.StatusCode, resp.Status)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ReadAll err %s", err)
	}

	return bodyBytes, nil
}
