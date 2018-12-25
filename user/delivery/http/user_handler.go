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
	Status    string `json:"status"`
}

type withdrawRateData struct {
	Rate   float32 `json:"rate"`
	Status string  `json:"status"`
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
	http.HandleFunc("/battle/",
		func(writer http.ResponseWriter, request *http.Request) {
			handler.generic_handler(writer, request, handler.battle_logic)
		})
	http.HandleFunc("/system/infrationTarget/",
		func(writer http.ResponseWriter, request *http.Request) {
			handler.generic_handler(writer, request, handler.infration_target_logic)
		})
	http.HandleFunc("/system/wallet/address/",
		func(writer http.ResponseWriter, request *http.Request) {
			handler.generic_handler(writer, request, handler.address_logic)
		})
	http.HandleFunc("/system/wallet/fixed_income/",
		func(writer http.ResponseWriter, request *http.Request) {
			handler.generic_handler(writer, request, handler.fixed_income_logic)
		})
	http.HandleFunc("/system/wallet/ratio_income/",
		func(writer http.ResponseWriter, request *http.Request) {
			handler.generic_handler(writer, request, handler.ratio_income_logic)
		})
	http.ListenAndServe(":8080", nil)
}

func (h *HttpUserHandler) generic_handler(writer http.ResponseWriter,
	request *http.Request, f func([]string) (interface{}, error)) {
	tmp := strings.SplitN(request.URL.Path, "/", 10)
	data, err := f(tmp)
	if err != nil {
		panic(err)
	}
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(writer).Encode(data)
}

func (h *HttpUserHandler) battle_logic(inputs []string) (interface{}, error) {
	if inputs[len(inputs)-1] != "" {
		username := inputs[len(inputs)-1]
		mony, earn, err := h.Ua.GetMony(username)
		if err != nil {
			return battleData{Status: err.Error()}, nil
		}
		return battleData{UserName: username, GotMony: earn, TotalMony: mony}, nil
	}
	return battleData{Status: fmt.Errorf("please input username").Error()}, nil
}

func (h *HttpUserHandler) infration_target_logic(inputs []string) (interface{}, error) {
	if inputs[len(inputs)-1] == "" {
		rate, err := h.Sa.WithdrawRate()
		if err != nil {
			return withdrawRateData{Status: err.Error()}, nil
		}
		return withdrawRateData{Rate: rate}, nil
	}
	return withdrawRateData{Status: fmt.Errorf("Does not support set").Error()}, nil
}

func (h *HttpUserHandler) address_logic(inputs []string) (interface{}, error) {
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

func (h *HttpUserHandler) fixed_income_logic(inputs []string) (interface{}, error) {
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

func (h *HttpUserHandler) ratio_income_logic(inputs []string) (interface{}, error) {
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
