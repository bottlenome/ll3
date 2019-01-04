package cript_currency

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/bottlenome/ll3/cript_currency"
	"net/http"
	"strconv"
)

type infraIOClient struct {
	endpoint string
}

func NewInfraIOClient(endpoint string) cript_currency.Client {
	return &infraIOClient{endpoint}
}

func (i *infraIOClient) httpPost(jsonStr string, data interface{}) error {
	req, err := http.NewRequest(
		"POST",
		i.endpoint,
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		return err
	}

	// Content-Type 設定
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(data)
}

type BalanceData struct {
	Id      uint64 `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Wei     string `json:"result"`
	Error   struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func (i *infraIOClient) Balance(address string) (uint64, error) {
	request := `{"jsonrpc":"2.0","method":"eth_getBalance","params":["` + address +
		`", "latest"],"id":1}`
	b := BalanceData{}
	err := i.httpPost(request, &b)
	if err != nil {
		return uint64(0), err
	}
	if b.Error.Code <= -32000 {
		return uint64(0), errors.New(b.Error.Message)
	}
	wei, err := strconv.ParseUint(b.Wei, 0, 64)
	if err != nil {
		return uint64(0), err
	}
	return wei, err
}
