package system

type SystemApplication interface {
	UpdateWithdrawRate() error
	WithdrawRate() (rate float32, err error)
	Wallet() (address string, err error)
	SetWallet(address string) error
	FixedIncome() (income float64, err error)
	SetFixedIncome(income float64) error
}
