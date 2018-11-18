package user

import (
	user "github.com/bottlenome/ll3/user"
)

type ll3UserApplication struct {
	repository user.UserRepository
}

func Newll3UserApplication(repository user.UserRepository) user.UserApplication {
	return &ll3UserApplication{repository}
}

func (l *ll3UserApplication) GetMony(username string, mony int64) (int64, error) {
	user, err := l.repository.GetByUserName(username)
	if err != nil {
		panic(err)
	}

	user.Mony += mony

	user, err = l.repository.Update(user)
	if err != nil {
		panic(err)
	}

	return user.Mony, err
}
