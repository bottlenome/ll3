package user

import (
	"errors"
	models "github.com/bottlenome/ll3/models"
	system "github.com/bottlenome/ll3/system"
	user "github.com/bottlenome/ll3/user"
	"testing"
)

type testRepository struct {
	user.UserRepository
	FakeGetByUserName func(username string) (*models.User, error)
	FakeUpdate        func(user *models.User) (*models.User, error)
	FakeCalcTotalMony func() (total float64, err error)
}

func (r *testRepository) GetByUserName(username string) (*models.User, error) {
	return r.FakeGetByUserName(username)
}

func (r *testRepository) Update(user *models.User) (*models.User, error) {
	return r.FakeUpdate(user)
}

func (r *testRepository) CalcTotalMony() (total float64, err error) {
	return r.FakeCalcTotalMony()
}

type testSystemRepository struct {
	system.SystemRepository
	FakeWithdrawRate func() (rate float32, err error)
}

func (s *testSystemRepository) WithdrawRate() (rate float32, err error) {
	return s.FakeWithdrawRate()
}

func TestGetMonyNormal(t *testing.T) {
	tr := &testRepository{
		FakeGetByUserName: func(username string) (*models.User, error) {
			u := models.User{
				UserName: username,
				Mony:     10,
			}
			return &u, nil
		},
		FakeUpdate: func(user *models.User) (*models.User, error) {
			return user, nil
		},
		FakeCalcTotalMony: func() (float64, error) {
			return 10.0, nil
		},
	}
	sr := &testSystemRepository{
		FakeWithdrawRate: func() (float32, error) {
			return 1.0, nil
		},
	}

	application := Newll3UserApplication(tr, sr)

	total, earn, err := application.GetMony("test_user")
	if err != nil {
		t.Fatalf("faild test %#v", err)
	}
	expect := uint64(21)
	if total != expect {
		t.Errorf("mismatch total:%d expect:%d", total, expect)
	}
	expect = uint64(11)
	if earn != expect {
		t.Errorf("mismatch earn:%d expect:%d", earn, expect)
	}
}

func TestGetMonyError(t *testing.T) {
	getByUserNameError := errors.New("GetByUserName")
	updateError := errors.New("Update")
	tr := &testRepository{
		FakeGetByUserName: func(username string) (*models.User, error) {
			return nil, getByUserNameError
		},
		FakeUpdate: func(user *models.User) (*models.User, error) {
			return nil, updateError
		},
	}
	withdrawRateError := errors.New("WithdrawRate")
	sr := &testSystemRepository{
		FakeWithdrawRate: func() (float32, error) {
			return 0.0, withdrawRateError
		},
	}

	application := Newll3UserApplication(tr, sr)
	_, _, err := application.GetMony("test_user")
	if err == nil {
		t.Errorf("does not occur error GetMony: %v", err)
	}

	tr.FakeGetByUserName = func(username string) (*models.User, error) {
		u := models.User{}
		return &u, nil
	}
	application = Newll3UserApplication(tr, sr)
	_, _, err = application.GetMony("test_user")
	if err == nil {
		t.Errorf("does not occur error GetMony: %v", withdrawRateError)
	}

	sr.FakeWithdrawRate = func() (float32, error) {
		return 0.1, nil
	}
	application = Newll3UserApplication(tr, sr)
	_, _, err = application.GetMony("test_user")
	if err == nil {
		t.Errorf("does not occur error update: %v", err)
	}
}
