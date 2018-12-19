package http

import (
	"encoding/json"
	"github.com/bottlenome/ll3/system"
	"github.com/bottlenome/ll3/user"
	"net/http"
	"strings"
)

type HttpUserHandler struct {
	Ua user.UserApplication
	Sa system.SystemApplication
}

type battleData struct {
	UserName  string `json:"userName"`
	GotMony   uint64 `json:"gotMony"`
	TotalMony uint64 `json:"totalMony"`
}

type withdrawRateData struct {
	Rate float32 `json:"rate"`
}

func NewUserHandler(ua user.UserApplication, sa system.SystemApplication) {
	handler := HttpUserHandler{
		Ua: ua,
		Sa: sa,
	}
	http.HandleFunc("/battle/", handler.battle)
	http.HandleFunc("/system/infrationTarget/", handler.infration_target)
	http.ListenAndServe(":8080", nil)
}

func (h *HttpUserHandler) battle(writer http.ResponseWriter, request *http.Request) {
	username := strings.SplitN(request.URL.Path, "/", 3)[2]

	// TODO change Ua. GetMony to return battleData
	mony, earn, err := h.Ua.GetMony(username)
	if err != nil {
		panic(err)
	}

	data := battleData{UserName: username, GotMony: earn, TotalMony: mony}
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(writer).Encode(data)
}

func (h *HttpUserHandler) infration_target(writer http.ResponseWriter, request *http.Request) {
	tmp := strings.SplitN(request.URL.Path, "/", 10)
	data := withdrawRateData{}
	if tmp[len(tmp)-1] == "" {
		// TODO change Sa.WithdrawRate() to return withdrawRateData
		rate, err := h.Sa.WithdrawRate()
		data.Rate = rate
		if err != nil {
			panic(err)
		}
	}

	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(writer).Encode(data)
}
