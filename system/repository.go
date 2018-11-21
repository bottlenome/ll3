package system

type SystemRepository interface {
	SetInflationTarget(target float32) error
	InflationTarget() (float32, error)
	SetUnit(unit uint64) error
	Unit() (uint64, error)
	SetRate(rate float32) error
	Rate() (float32, error)
	SetWithdrawRate(rate float32) error
}
