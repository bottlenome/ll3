package http

import (
	"encoding/json"
	"fmt"
	"github.com/bottlenome/ll3/system"
	"github.com/bottlenome/ll3/user"
	"net/http"
	"strconv"
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

type addressData struct {
	Address string `json:"address"`
	Status  string `json:"status"`
}

type fixedIncomeData struct {
	FixedIncome float64 `json:"fixed_income"`
	Status      string  `json:"status"`
}

type ratioIncomeData struct {
	RatioIncome float64 `json:"ratio_income"`
	Status      string  `json:"status"`
}

func NewUserHandler(ua user.UserApplication, sa system.SystemApplication) {
	handler := HttpUserHandler{
		Ua: ua,
		Sa: sa,
	}
	http.HandleFunc("/battle/", handler.battle)
	http.HandleFunc("/system/infrationTarget/", handler.infration_target)
	http.HandleFunc("/system/wallet/address/", handler.address)
	http.HandleFunc("/system/wallet/fixed_income/", handler.fixed_income)
	http.HandleFunc("/system/wallet/ratio_income/", handler.ratio_income)
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

func (h *HttpUserHandler) address_logic(inputs []string) (addressData, error) {
	address := inputs[len(inputs)-1]
	if address != "" {
		err := h.Sa.SetWallet(address)
		if err != nil {
			return addressData{Address: "",
				Status: err.Error()}, nil
		}
		return addressData{Address: address}, nil
	} else {
		address, err := h.Sa.Wallet()
		if err != nil {
			return addressData{Address: "",
				Status: err.Error()}, nil
		}
		return addressData{Address: address}, nil
	}
}

func (h *HttpUserHandler) address(writer http.ResponseWriter, request *http.Request) {
	tmp := strings.SplitN(request.URL.Path, "/", 10)
	data, err := h.address_logic(tmp)
	if err != nil {
		panic(err)
	}
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(writer).Encode(data)
}

func (h *HttpUserHandler) fixed_income_logic(inputs []string) (fixedIncomeData, error) {
	income := inputs[len(inputs)-1]
	if income != "" {
		income64, err := strconv.ParseFloat(income, 64)
		if err != nil {
			return fixedIncomeData{FixedIncome: 0.0,
				Status: fmt.Errorf("invalid format: %v", income).Error()}, nil
		}
		err = h.Sa.SetFixedIncome(income64)
		if err != nil {
			return fixedIncomeData{FixedIncome: 0.0,
				Status: err.Error()}, nil
		}
		return fixedIncomeData{FixedIncome: income64}, nil
	} else {
		income, err := h.Sa.FixedIncome()
		if err != nil {
			return fixedIncomeData{FixedIncome: 0.0,
				Status: err.Error()}, nil
		}
		return fixedIncomeData{FixedIncome: income}, nil
	}
}

func (h *HttpUserHandler) fixed_income(writer http.ResponseWriter, request *http.Request) {
	tmp := strings.SplitN(request.URL.Path, "/", 10)
	data, err := h.fixed_income_logic(tmp)
	if err != nil {
		panic(err)
	}
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(writer).Encode(data)
}

func (h *HttpUserHandler) ratio_income_logic(inputs []string) (ratioIncomeData, error) {
	income := inputs[len(inputs)-1]
	if income != "" {
		income64, err := strconv.ParseFloat(income, 64)
		if err != nil {
			return ratioIncomeData{RatioIncome: 0.0,
				Status: fmt.Errorf("invalid format: %v", income).Error()}, nil
		}
		err = h.Sa.SetRatioIncome(income64)
		if err != nil {
			return ratioIncomeData{RatioIncome: 0.0,
				Status: err.Error()}, nil
		}
		return ratioIncomeData{RatioIncome: income64}, nil
	} else {
		income, err := h.Sa.RatioIncome()
		if err != nil {
			return ratioIncomeData{RatioIncome: 0.0,
				Status: err.Error()}, nil
		}
		return ratioIncomeData{RatioIncome: income}, nil
	}
}

func (h *HttpUserHandler) ratio_income(writer http.ResponseWriter, request *http.Request) {
	tmp := strings.SplitN(request.URL.Path, "/", 10)
	data, err := h.ratio_income_logic(tmp)
	if err != nil {
		panic(err)
	}
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(writer).Encode(data)
}
