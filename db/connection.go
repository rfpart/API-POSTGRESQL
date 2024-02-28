package db // este pacote é responsável pela conexão com banco de dados

import (
	"database/sql"
	"fmt"

	configs "github.com/API-POSTGRESQL/config"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) { // esta função abre a conexão com banco de dados
	conf := configs.GetDb() // a variável conf recebe as configurações do banco de dados para poder conectar

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	conn, err := sql.Open("postgres", sc)
	if err != nil {
		panic(err)
	}

	err = conn.Ping() //

	return conn, err
}
