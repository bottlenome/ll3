package http

import (
	"encoding/json"
	"github.com/bottlenome/ll3/user"
	"net/http"
	"strings"
)

type HttpUserHandler struct {
	Ua user.UserApplication
}

type battleData struct {
	UserName  string `json:"userName"`
	GotMony   uint64 `json:"gotMony"`
	TotalMony uint64 `json:"totalMony"`
}

func NewUserHandler(ua user.UserApplication) {
	handler := HttpUserHandler{
		Ua: ua,
	}
	http.HandleFunc("/battle/", handler.battle)
	http.ListenAndServe(":8080", nil)
}

func (h *HttpUserHandler) battle(writer http.ResponseWriter, request *http.Request) {
	username := strings.SplitN(request.URL.Path, "/", 3)[2]

	mony, earn, err := h.Ua.GetMony(username)
	if err != nil {
		panic(err)
	}

	data := battleData{UserName: username, GotMony: earn, TotalMony: mony}
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(writer).Encode(data)
}
