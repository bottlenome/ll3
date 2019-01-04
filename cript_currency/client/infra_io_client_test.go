package cript_currency

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if os.Getenv("INFRA_IO_ENDPOINT") == "" {
		println("Load local environment value")
		err := godotenv.Load(os.ExpandEnv("${GOPATH}/src/github.com/bottlenome/ll3/.env"))
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	code := m.Run()

	os.Exit(code)
}

func balanceHandler(result string, code int) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		type data struct {
			Id      uint64 `json:"id"`
			Jsonrpc string `json:"jsonrpc"`
			Error   struct {
				code int `json:"code"`
			} `json:"error"`
		}
		var val data
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return
		}
		err = json.Unmarshal(b, &val)
		if err != nil {
			return
		}
		jsonStr := fmt.Sprintf(`{
"id":%d,
"jsonrpc":"%s",
"result":"%s",
"error":{
  "code":%d
  }
}`, val.Id, val.Jsonrpc, result, code)
		w.Write([]byte(jsonStr))
	}
}

func TestBalanceNormal(t *testing.T) {
	url := ""
	if os.Getenv("ENABLE_HTTP") == "ON" {
		url = os.Getenv("INFRA_IO_ENDPOINT")
	} else {
		testServer := httptest.NewServer(http.HandlerFunc(balanceHandler("0x6f05b59d3b20000", 0)))
		url = testServer.URL
		defer testServer.Close()
	}
	ii := NewInfraIOClient(url)
	balance, err := ii.Balance(os.Getenv("TEST_ADMIN_WALLET"))
	if err != nil {
		t.Errorf("Blaance error: %v", err)
	}
	expect := uint64(0x6f05b59d3b20000)
	if balance != expect {
		t.Errorf("balance mismatch: %d expect %d", balance, expect)
	}
}

func TestBalanceError(t *testing.T) {
	url := ""
	ii := NewInfraIOClient(url)
	wei, _ := ii.Balance("hoge")
	if wei != uint64(0) {
		t.Errorf("Some error happen at Blance error")
	}

	testServer := httptest.NewServer(http.HandlerFunc(balanceHandler("", -32000)))
	url = testServer.URL
	ii = NewInfraIOClient(url)
	wei, _ = ii.Balance(os.Getenv("TEST_ADMIN_WALLET"))
	if wei != uint64(0) {
		t.Errorf("Some error happen at Blance error")
	}
	testServer.Close()

	testServer = httptest.NewServer(http.HandlerFunc(balanceHandler("hoge", 0)))
	url = testServer.URL
	ii = NewInfraIOClient(url)
	wei, _ = ii.Balance(os.Getenv("TEST_ADMIN_WALLET"))
	if wei != uint64(0) {
		t.Errorf("Some error happen at Blance error")
	}
	testServer.Close()
}
