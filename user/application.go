package user

type UserApplication interface {
	GetMony(username string) (uint64, uint64, error)
}
