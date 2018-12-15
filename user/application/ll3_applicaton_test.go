package user

import (
	models "github.com/bottlenome/ll3/models"
	system "github.com/bottlenome/ll3/system"
	"testing"
)

type testRepository struct {
}

func (r *testRepository) GetByUserName(username string) (*models.User, error) {
	u := models.User{
		UserName: username,
		Mony:     10,
	}
	return &u, nil
}

func (r *testRepository) Update(user *models.User) (*models.User, error) {
	u := models.User{
		UserName: user.UserName,
		Mony:     user.Mony,
	}
	return &u, nil
}

type testSystemRepository struct {
	system.SystemRepository
}

func (s *testSystemRepository) WithdrawRate() (rate float32, err error) {
	return 1.0, nil
}

func (r *testRepository) CalcTotalMony() (total float64, err error) {
	return 10.0, nil
}

func TestGetMonyNormal(t *testing.T) {
	tr := &testRepository{}
	sr := &testSystemRepository{}

	application := &ll3UserApplication{
		repository:       tr,
		systemRepository: sr,
	}
	total, earn, err := application.GetMony("test_user")
	if err != nil {
		t.Fatalf("faild test %#v", err)
	}
	if total != 21 {
		t.Errorf("mismatch total:%d expect:%d", 5, total)
	}
	if earn != 11 {
		t.Errorf("mismatch earn:%d expect:%d", 5, earn)
	}
}
