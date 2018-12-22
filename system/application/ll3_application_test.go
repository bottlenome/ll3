package system

import (
	"errors"
	system "github.com/bottlenome/ll3/system"
	user "github.com/bottlenome/ll3/user"
	"testing"
)

type fakeSystemRepository struct {
	system.SystemRepository
	FakeRate            func() (float32, error)
	FakeUnit            func() (uint64, error)
	FakeSetWithdrawRate func(rate float32) error
}

type fakeUserRepository struct {
	user.UserRepository
	FakeCalcTotalMony func() (float64, error)
}

func (f *fakeSystemRepository) Rate() (float32, error) {
	return f.FakeRate()
}

func (f *fakeSystemRepository) Unit() (uint64, error) {
	return f.FakeUnit()
}

func (f *fakeSystemRepository) SetWithdrawRate(rate float32) error {
	return f.FakeSetWithdrawRate(rate)
}

func (f *fakeUserRepository) CalcTotalMony() (float64, error) {
	return f.FakeCalcTotalMony()
}

func TestWithdrawRateNormal(t *testing.T) {
	sr := fakeSystemRepository{
		FakeRate: func() (float32, error) {
			return 2.0, nil
		},
	}
	ur := fakeUserRepository{}
	// TODO(bottlenome) I don't know why &sr works fine
	application := Newll3SystemApplication(&sr, &ur)
	rate, err := application.WithdrawRate()
	if err != nil {
		t.Fatalf("faild test %#v", err)
	}
	expect := float32(2.0)
	if rate != expect {
		t.Errorf("mismatch rate: %f expect %f", rate, expect)
	}
}

func TestWithdrawRateError(t *testing.T) {
	sr := fakeSystemRepository{
		FakeRate: func() (float32, error) {
			return 0.0, errors.New("Rate")
		},
	}
	ur := fakeUserRepository{}
	// TODO(bottlenome) I don't know why &sr works fine
	application := Newll3SystemApplication(&sr, &ur)
	_, err := application.WithdrawRate()
	if err == nil {
		t.Fatalf("faild test %#v", err)
	}
}

func TestUpdateWithrawRateNormal(t *testing.T) {
	sr := fakeSystemRepository{
		FakeRate: func() (float32, error) {
			return 2.0, nil
		},
		FakeUnit: func() (uint64, error) {
			return 3000, nil
		},
		FakeSetWithdrawRate: func(rate float32) error {
			return nil
		},
	}
	ur := fakeUserRepository{
		FakeCalcTotalMony: func() (float64, error) {
			return 10001, nil
		},
	}
	// TODO(bottlenome) I don't know why &sr works fine
	application := Newll3SystemApplication(&sr, &ur)
	err := application.UpdateWithdrawRate()
	if err != nil {
		t.Fatalf("faild test %#v", err)
	}
}

func TestUpdateWithrawRateError(t *testing.T) {
	sr := fakeSystemRepository{
		FakeRate: func() (float32, error) {
			return 0.0, errors.New("Rate")
		},
	}
	ur := fakeUserRepository{}
	// TODO(bottlenome) I don't know why &sr works fine
	application := Newll3SystemApplication(&sr, &ur)
	err := application.UpdateWithdrawRate()
	if err == nil {
		t.Fatalf("faild test %#v", err)
	}

	sr = fakeSystemRepository{
		FakeRate: func() (float32, error) {
			return 2.0, nil
		},
		FakeUnit: func() (uint64, error) {
			return 0, errors.New("Unit")
		},
	}
	ur = fakeUserRepository{}
	// TODO(bottlenome) I don't know why &sr works fine
	application = Newll3SystemApplication(&sr, &ur)
	err = application.UpdateWithdrawRate()
	if err == nil {
		t.Fatalf("faild test %#v", err)
	}

	sr = fakeSystemRepository{
		FakeRate: func() (float32, error) {
			return 2.0, nil
		},
		FakeUnit: func() (uint64, error) {
			return 3000, nil
		},
		FakeSetWithdrawRate: func(rate float32) error {
			return nil
		},
	}
	ur = fakeUserRepository{
		FakeCalcTotalMony: func() (float64, error) {
			return 0, errors.New("CalcTotalMony")
		},
	}
	// TODO(bottlenome) I don't know why &sr works fine
	application = Newll3SystemApplication(&sr, &ur)
	err = application.UpdateWithdrawRate()
	if err == nil {
		t.Fatalf("faild test %#v", err)
	}

	sr = fakeSystemRepository{
		FakeRate: func() (float32, error) {
			return 2.0, nil
		},
		FakeUnit: func() (uint64, error) {
			return 3000, nil
		},
		FakeSetWithdrawRate: func(rate float32) error {
			return errors.New("SetWithdrawRate")
		},
	}
	ur = fakeUserRepository{
		FakeCalcTotalMony: func() (float64, error) {
			return 10001, nil
		},
	}
	// TODO(bottlenome) I don't know why &sr works fine
	application = Newll3SystemApplication(&sr, &ur)
	err = application.UpdateWithdrawRate()
	if err == nil {
		t.Fatalf("faild test %#v", err)
	}
}
