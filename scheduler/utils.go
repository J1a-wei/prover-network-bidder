package scheduler

import (
	"fmt"
	"io"
	"net/http"
	ulib "net/url"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func DownloadFile(url string) ([]byte, error) {
	if strings.HasPrefix(url, "file://") {
		u, err := ulib.Parse(url)
		if err != nil {
			return nil, fmt.Errorf("invalid file url %v", err)
		}
		raw, err := os.ReadFile(u.Path)
		if err != nil {
			return nil, fmt.Errorf("open file %s err: %v", u.Path, err)
		}
		return raw, nil
	}
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

type JsonError struct {
	Code    int
	Message string
	Data    string
}

func ParseSolCustomErrorName(contractABI string, errData []byte) (string, error) {
	if len(errData) < 4 {
		return "", fmt.Errorf("invalid errData")
	}

	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		return "", fmt.Errorf("abi.JSON err: %s", err)
	}

	for _, errDef := range parsedABI.Errors {
		if common.Bytes2Hex(errData[:4]) == common.Bytes2Hex(errDef.ID[:4]) {
			return errDef.Name, nil
		}
	}

	return "", nil
}
