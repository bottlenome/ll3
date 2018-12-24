package system

import (
	"fmt"
	"github.com/bottlenome/ll3/system"
	"github.com/bottlenome/ll3/user"
)

type ll3SystemApplication struct {
	repository     system.SystemRepository
	userRepository user.UserRepository
}

func Newll3SystemApplication(
	repository system.SystemRepository,
	userRepository user.UserRepository) system.SystemApplication {
	return &ll3SystemApplication{repository, userRepository}
}

func (l *ll3SystemApplication) getRate() (float32, error) {
	// check rate
	rate, err := l.repository.Rate()
	if err != nil {
		return 0.0, fmt.Errorf("faild to get rate : %v", err)
	}
	return rate, nil
}

func (l *ll3SystemApplication) WithdrawRate() (float32, error) {
	return l.getRate()
}

func (l *ll3SystemApplication) UpdateWithdrawRate() error {
	// check rate
	rate, err := l.getRate()
	if err != nil {
		return err
	}
	unit, err := l.repository.Unit()
	if err != nil {
		return fmt.Errorf("faild to get unit : %v", err)
	}
	target := float64(rate) * float64(unit)
	// check total mony
	total, err := l.userRepository.CalcTotalMony()
	if err != nil {
		return fmt.Errorf("faild to CalcTotalMony : %v", err)
	}
	// modify withdraw_rate
	err = l.repository.SetWithdrawRate(float32((target - total) / 10000.0))
	return err
}

func (l *ll3SystemApplication) Wallet() (string, error) {
	return l.repository.Wallet()
}

func (l *ll3SystemApplication) SetWallet(address string) error {
	if len(address) != 42 {
		return fmt.Errorf("invalid address size : %d", len(address))
	}
	return l.repository.SetWallet(address)
}

func (l *ll3SystemApplication) FixedIncome() (float64, error) {
	return l.repository.FixedIncome()
}

func (l *ll3SystemApplication) SetFixedIncome(income float64) error {
	return l.repository.SetFixedIncome(income)
}
