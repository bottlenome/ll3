package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strings"
)

func env_load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loding .env file")
	}
}

func db_test() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/ll3")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}

	var (
		username string
		mony     int
	)

	for rows.Next() {
		err := rows.Scan(&username, &mony)
		if err != nil {
			panic(err)
		}
		fmt.Println(username, mony)
	}

	err = db.QueryRow("SELECT mony FROM users WHERE username=?", "test_user").Scan(&mony)
	if err != nil {
		panic(err)
	}
}

func main() {
	env_load()
	db_test()

	http.HandleFunc("/battle/", battle)

	http.HandleFunc("/hello", hello)

	http.HandleFunc("/weather/", func(writer http.ResponseWriter, request *http.Request) {
		city := strings.SplitN(request.URL.Path, "/", 3)[2]

		data, err := query(city)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(writer).Encode(data)
	})

	http.ListenAndServe(":8080", nil)
}

// Web APIs
func battle(writer http.ResponseWriter, request *http.Request) {
	EARN := int64(5)
	fmt.Println(request.URL.Path)
	username := strings.SplitN(request.URL.Path, "/", 3)[2]
	fmt.Println(username)

	//
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/ll3")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("SELECT * FROM users WHERE username = ?")
	if err != nil {
		panic(err)
	}

	var mony int64

	err = db.QueryRow("SELECT mony FROM users WHERE username=?", "test_user").Scan(&mony)
	if err != nil {
		panic(err)
	}

	mony += EARN

	stmt, err = db.Prepare("UPDATE users SET mony=? WHERE username=?")
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(mony, username)
	if err != nil {
		panic(err)
	}
	rowCount, err := res.RowsAffected()
	if rowCount != 1 {
		panic(res)
	}

	data := battleData{UserName: username, GotMony: EARN, TotalMony: mony}
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(writer).Encode(data)
}

type battleData struct {
	UserName  string `json:"userName"`
	GotMony   int64  `json:"gotMony"`
	TotalMony int64  `json:"totalMony"`
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}

func query(city string) (weatherData, error) {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city +
		"&APPID=" + os.Getenv("OPEN_WEATHER_MAP_KEY"))
	if err != nil {
		return weatherData{}, err
	}

	defer resp.Body.Close()

	var data weatherData

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return weatherData{}, err
	}

	return data, nil
}

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

// Db Access
