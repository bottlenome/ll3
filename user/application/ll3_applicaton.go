package user

import (
	"github.com/bottlenome/ll3/system"
	"github.com/bottlenome/ll3/user"
)

type ll3UserApplication struct {
	repository       user.UserRepository
	systemRepository system.SystemRepository
}

func Newll3UserApplication(
	repository user.UserRepository,
	systemRepository system.SystemRepository) user.UserApplication {
	return &ll3UserApplication{repository, systemRepository}
}

func (l *ll3UserApplication) GetMony(username string) (uint64, uint64, error) {
	user, err := l.repository.GetByUserName(username)
	if err != nil {
		panic(err)
	}

	withdrawRate, err := l.systemRepository.WithdrawRate()
	earn := uint64(10*withdrawRate) + 1
	user.Mony += int64(earn)

	user, err = l.repository.Update(user)
	if err != nil {
		panic(err)
	}

	return uint64(user.Mony), earn, err
}
