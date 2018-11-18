package user

import (
	models "github.com/bottlenome/ll3/models"
	"testing"
)

type testRepository struct {
}

func (r *testRepository) GetByUserName(username string) (*models.User, error) {
	u := models.User{
		UserName: username,
		Mony:     0,
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

func TestGetMonyNormal(t *testing.T) {
	tr := &testRepository{}

	application := &ll3UserApplication{
		repository: tr,
	}
	result, err := application.GetMony("test_user", 5)
	if err != nil {
		t.Fatalf("faild test %#v", err)
	}
	if result != 5 {
		t.Fatal("failed test")
	}
}
