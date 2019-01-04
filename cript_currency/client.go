package cript_currency

type Client interface {
	Balance(address string) (uint64, error)
}
