package db

import (
	"database/sql"
	"fmt"

	"github.com/aprendagolang/api-pgsql/configs"
)

func OpenConnection() {
	conf := configs.GetDb()

	sc := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	conn, err := sql.Open("postgres", sc)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}
