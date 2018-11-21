package main

import (
	"database/sql"
	"fmt"
	systemRepository "github.com/bottlenome/ll3/system/repository"
	userApplication "github.com/bottlenome/ll3/user/application"
	httpDeliver "github.com/bottlenome/ll3/user/delivery/http"
	userRepository "github.com/bottlenome/ll3/user/repository"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
)

func env_load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loding .env file")
	}
}

func main() {
	env_load()

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/ll3")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	mysql_user := userRepository.NewMysqlUserRepository(db)
	ll3_application := userApplication.Newll3UserApplication(mysql_user)

	// test
	mysql_system := systemRepository.NewMysqlSystemRepository(db)
	fmt.Println(mysql_system.SetInflationTarget(1.02))
	fmt.Println(mysql_system.InflationTarget())
	fmt.Println(mysql_system.SetUnit(1000))
	fmt.Println(mysql_system.Unit())
	fmt.Println(mysql_system.SetRate(10))
	fmt.Println(mysql_system.Rate())
	fmt.Println(mysql_system.SetWithdrawRate(10))

	httpDeliver.NewUserHandler(ll3_application)
}
