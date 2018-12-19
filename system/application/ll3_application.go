package system

import (
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

func (l *ll3SystemApplication) WithdrawRate() (float32, error) {
	// check rate
	rate, err := l.repository.Rate()
	if err != nil {
		panic(err)
	}
	return rate, nil
}
func (l *ll3SystemApplication) UpdateWithdrawRate() error {
	// check rate
	rate, err := l.repository.Rate()
	if err != nil {
		panic(err)
	}
	unit, err := l.repository.Unit()
	if err != nil {
		panic(err)
	}
	target := float64(rate) * float64(unit)
	// check total mony
	total, err := l.userRepository.CalcTotalMony()
	if err != nil {
		panic(err)
	}
	// modify withdraw_rate
	err = l.repository.SetWithdrawRate(float32((target - total) / 10000.0))
	return err
}
