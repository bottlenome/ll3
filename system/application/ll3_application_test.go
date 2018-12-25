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
	FakeSetWallet       func(address string) error
	FakeWallet          func() (string, error)
	FakeSetFixedIncome  func(income float64) error
	FakeFixedIncome     func() (float64, error)
	FakeSetRatioIncome  func(income float64) error
	FakeRatioIncome     func() (float64, error)
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

func (f *fakeSystemRepository) SetWallet(address string) error {
	return f.FakeSetWallet(address)
}

func (f *fakeSystemRepository) Wallet() (string, error) {
	return f.FakeWallet()
}

func (f *fakeSystemRepository) SetFixedIncome(income float64) error {
	return f.FakeSetFixedIncome(income)
}

func (f *fakeSystemRepository) FixedIncome() (float64, error) {
	return f.FakeFixedIncome()
}

func (f *fakeSystemRepository) SetRatioIncome(income float64) error {
	return f.FakeSetRatioIncome(income)
}

func (f *fakeSystemRepository) RatioIncome() (float64, error) {
	return f.FakeRatioIncome()
}

// user.repository
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

func TestSetWalletNormal(t *testing.T) {
	sr := fakeSystemRepository{
		FakeSetWallet: func(address string) error {
			return nil
		},
	}
	ur := fakeUserRepository{}
	// TODO(bottlenome) I don't know why &sr works fine
	application := Newll3SystemApplication(&sr, &ur)
	err := application.SetWallet("012345678901234567890123456789012345678901")
	if err != nil {
		t.Fatalf("faild test %#v", err)
	}
}

func TestSetWalletError(t *testing.T) {
	sr := fakeSystemRepository{}
	ur := fakeUserRepository{}
	// TODO(bottlenome) I don't know why &sr works fine
	application := Newll3SystemApplication(&sr, &ur)
	err := application.SetWallet("")
	if err == nil {
		t.Fatalf("faild test %#v", err)
	}
}

func TestWalletNormal(t *testing.T) {
	sr := fakeSystemRepository{
		FakeWallet: func() (string, error) {
			return "hoge", nil
		},
	}
	ur := fakeUserRepository{}
	// TODO(bottlenome) I don't know why &sr works fine
	application := Newll3SystemApplication(&sr, &ur)
	address, err := application.Wallet()
	if err != nil || address != "hoge" {
		t.Fatalf("faild test %#v", err)
	}
}

func TestSetFixedIncomeNormal(t *testing.T) {
	sr := fakeSystemRepository{
		FakeSetFixedIncome: func(float64) error {
			return nil
		},
	}
	ur := fakeUserRepository{}
	// TODO(bottlenome) I don't know why &sr works fine
	application := Newll3SystemApplication(&sr, &ur)
	err := application.SetFixedIncome(float64(10.0))
	if err != nil {
		t.Fatalf("faild test %#v", err)
	}
}

func TestSetFixedIncomeError(t *testing.T) {
	sr := fakeSystemRepository{
		FakeSetFixedIncome: func(float64) error {
			return errors.New("SetFixedIncome")
		},
	}
	ur := fakeUserRepository{}
	// TODO(bottlenome) I don't know why &sr works fine
	application := Newll3SystemApplication(&sr, &ur)
	err := application.SetFixedIncome(float64(10.0))
	if err == nil {
		t.Fatalf("faild test %#v", err)
	}
}

func TestFixedIncomeNormal(t *testing.T) {
	sr := fakeSystemRepository{
		FakeFixedIncome: func() (float64, error) {
			return float64(10.0), nil
		},
	}
	ur := fakeUserRepository{}
	// TODO(bottlenome) I don't know why &sr works fine
	application := Newll3SystemApplication(&sr, &ur)
	income, err := application.FixedIncome()
	if err != nil || income != 10.0 {
		t.Fatalf("faild test %#v", err)
	}
}

func TestFixedIncomeError(t *testing.T) {
	sr := fakeSystemRepository{
		FakeFixedIncome: func() (float64, error) {
			return float64(10.0), errors.New("FixedIncome")
		},
	}
	ur := fakeUserRepository{}
	// TODO(bottlenome) I don't know why &sr works fine
	application := Newll3SystemApplication(&sr, &ur)
	_, err := application.FixedIncome()
	if err == nil {
		t.Fatalf("faild test %#v", err)
	}
}

func TestSetRatioIncomeNormal(t *testing.T) {
	sr := fakeSystemRepository{
		FakeSetRatioIncome: func(float64) error {
			return nil
		},
	}
	ur := fakeUserRepository{}
	// TODO(bottlenome) I don't know why &sr works fine
	application := Newll3SystemApplication(&sr, &ur)
	err := application.SetRatioIncome(float64(0.1))
	if err != nil {
		t.Fatalf("faild test %#v", err)
	}
}

func TestSetRatioIncomeError(t *testing.T) {
	sr := fakeSystemRepository{
		FakeSetRatioIncome: func(float64) error {
			return errors.New("SetRatioIncome")
		},
	}
	ur := fakeUserRepository{}
	// TODO(bottlenome) I don't know why &sr works fine
	application := Newll3SystemApplication(&sr, &ur)
	err := application.SetRatioIncome(float64(10.0))
	if err == nil {
		t.Fatalf("faild test %#v", err)
	}
	err = application.SetRatioIncome(float64(0.1))
	if err == nil {
		t.Fatalf("faild test %#v", err)
	}
}

func TestRatioIncomeNormal(t *testing.T) {
	sr := fakeSystemRepository{
		FakeRatioIncome: func() (float64, error) {
			return float64(0.1), nil
		},
	}
	ur := fakeUserRepository{}
	// TODO(bottlenome) I don't know why &sr works fine
	application := Newll3SystemApplication(&sr, &ur)
	income, err := application.RatioIncome()
	if err != nil || income != 0.1 {
		t.Fatalf("faild test %#v", err)
	}
}

func TestRatioIncomeError(t *testing.T) {
	sr := fakeSystemRepository{
		FakeRatioIncome: func() (float64, error) {
			return float64(0.1), errors.New("RatioIncome")
		},
	}
	ur := fakeUserRepository{}
	// TODO(bottlenome) I don't know why &sr works fine
	application := Newll3SystemApplication(&sr, &ur)
	_, err := application.RatioIncome()
	if err == nil {
		t.Fatalf("faild test %#v", err)
	}
}
