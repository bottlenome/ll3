package system

type SystemApplication interface {
	UpdateWithdrawRate() error
	WithdrawRate() (rate float32, err error)
}
