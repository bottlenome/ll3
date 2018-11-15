package repository

import (
	"database/sql"
	models "github.com/bottlenome/ll3/models"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlUserRepository struct{}

func (MysqlUserRepository) GetByUserName(username string) (*models.User, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/ll3")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var user models.User

	err = db.QueryRow("SELECT * FROM users WHERE username=?", username).
		Scan(&user.UserName,
			&user.Mony)
	if err != nil {
		panic(err)
	}

	return &user, err
}

func (MysqlUserRepository) Update(user *models.User) (*models.User, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/ll3")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE users SET mony=? WHERE username=?")
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(user.Mony, user.UserName)
	if err != nil {
		panic(err)
	}
	rowCount, err := res.RowsAffected()
	if rowCount != 1 {
		panic(res)
	}

	return user, err
}
