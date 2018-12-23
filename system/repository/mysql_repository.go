package system

import (
	"database/sql"
	"fmt"
	system "github.com/bottlenome/ll3/system"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type mysqlSystemRepository struct {
	db *sql.DB
}

func NewMysqlSystemRepository(db *sql.DB) system.SystemRepository {
	return &mysqlSystemRepository{db}
}

func (m *mysqlSystemRepository) set(target string, value string) error {
	res, err := m.db.Exec("UPDATE mony SET "+target+"=? WHERE type=?", value, "wei")
	if err != nil {
		panic(err)
	}
	rowCount, err := res.RowsAffected()
	if rowCount == 0 {
		log.Print("Update for " + target + " has no effect")
		return nil
	}

	return err
}

func (m *mysqlSystemRepository) get(target string) (float32, error) {
	value := float32(0.0)
	err := m.db.QueryRow("SELECT "+target+" FROM mony WHERE type=?", "wei").
		Scan(&value)
	if err != nil {
		panic(err)
	}
	return value, err
}

func (m *mysqlSystemRepository) getUint64(target string) (uint64, error) {
	value := uint64(0.0)
	err := m.db.QueryRow("SELECT "+target+" FROM mony WHERE type=?", "wei").
		Scan(&value)
	if err != nil {
		panic(err)
	}
	return value, err
}

func (m *mysqlSystemRepository) getString(target string) (string, error) {
	value := string("")
	err := m.db.QueryRow("SELECT "+target+" FROM mony WHERE type=?", "wei").
		Scan(&value)
	if err != nil {
		panic(err)
	}
	return value, err
}

func (m *mysqlSystemRepository) SetInflationTarget(target float32) error {
	return m.set("inflation_target", fmt.Sprintf("%f", target))
}

func (m *mysqlSystemRepository) InflationTarget() (float32, error) {
	return m.get("inflation_target")
}

func (m *mysqlSystemRepository) SetUnit(unit uint64) error {
	return m.set("unit", fmt.Sprintf("%d", unit))
}

func (m *mysqlSystemRepository) Unit() (uint64, error) {
	return m.getUint64("unit")
}

func (m *mysqlSystemRepository) SetRate(rate float32) error {
	return m.set("rate", fmt.Sprintf("%f", rate))
}

func (m *mysqlSystemRepository) Rate() (float32, error) {
	return m.get("rate")
}

func (m *mysqlSystemRepository) SetWithdrawRate(rate float32) error {
	return m.set("withdraw_rate", fmt.Sprintf("%f", rate))
}

func (m *mysqlSystemRepository) WithdrawRate() (float32, error) {
	return m.get("withdraw_rate")
}

func (m *mysqlSystemRepository) SetWallet(address string) error {
	return m.set("wallet_address", address)
}

func (m *mysqlSystemRepository) Wallet() (string, error) {
	return m.getString("wallet_address")
}
