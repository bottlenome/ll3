package user

import (
	"database/sql"
	models "github.com/bottlenome/ll3/models"
	user "github.com/bottlenome/ll3/user"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type mysqlUserRepository struct {
	db *sql.DB
}

func NewMysqlUserRepository(db *sql.DB) user.UserRepository {
	return &mysqlUserRepository{db}
}

func (m *mysqlUserRepository) GetByUserName(username string) (*models.User, error) {
	user := models.User{}
	err := m.db.QueryRow("SELECT * FROM users WHERE username=?", username).
		Scan(&user.UserName,
			&user.Mony)
	if err != nil {
		panic(err)
	}

	return &user, err
}

func (m *mysqlUserRepository) Update(user *models.User) (*models.User, error) {
	stmt, err := m.db.Prepare("UPDATE users SET mony=? WHERE username=?")
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(user.Mony, user.UserName)
	if err != nil {
		panic(err)
	}
	rowCount, err := res.RowsAffected()
	if rowCount == 0 {
		log.Print("Update for " + user.UserName + " has no effect")
		return user, nil
	}

	return user, err
}

func (m *mysqlUserRepository) CalcTotalMony() (total float64, err error) {
	total = 0.0
	err = m.db.QueryRow("SELECT sum(mony) FROM users").Scan(&total)
	if err != nil {
		panic(err)
	}
	return total, err
}
