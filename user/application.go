package user

type UserApplication interface {
	GetMony(username string, mony int64) (int64, error)
}
