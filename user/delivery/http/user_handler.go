package http

import (
	"encoding/json"
	"github.com/bottlenome/ll3/user"
	"net/http"
	"strings"
)

type HttpUserHandler struct {
	Ur *user.UserRepository
}

type battleData struct {
	UserName  string `json:"userName"`
	GotMony   int64  `json:"gotMony"`
	TotalMony int64  `json:"totalMony"`
}

func NewUserHandler(ur user.UserRepository) {
	handler := HttpUserHandler{
		Ur: ur,
	}
	http.HandleFunc("/battle/", handler.battle)
	http.ListenAndServe(":8080", nil)
}

func (h *HttpUserHandler) battle(writer http.ResponseWriter, request *http.Request) {
	username := strings.SplitN(request.URL.Path, "/", 3)[2]

	user, err := h.Ur.GetByUserName(username)
	if err != nil {
		panic(err)
	}

	var mony int64
	EARN := int64(5)

	user.Mony += 5

	user, err = h.Ur.Update(user)
	if err != nil {
		panic(err)
	}

	data := battleData{UserName: username, GotMony: EARN, TotalMony: user.mony}
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(writer).Encode(data)
}
