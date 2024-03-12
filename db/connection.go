package db // este pacote é responsável pela conexão com banco de dados

import (
	"database/sql" //pacote usado em go para se conectar ao PostgreSQL e excutar consultas
	"fmt"

	"github.com/API-POSTGRESQL/configs" // importa o pacote de configuração do banco de dados
	_ "github.com/lib/pq"               //Driver que permite conectar aplicativos go ao banco de dados PortgreSQL
)

func OpenConnection() (*sql.DB, error) { // esta função abre a conexão com banco de dados
	conf := configs.GetDb() // a variável conf recebe as configurações do banco de dados para poder conectar

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database) // cria uma string de conexão com os dados armazenados na struct DBConfig que está no pakage config

	conn, err := sql.Open("postgres", sc) // sql.Open() abre uma nova conexão com o banco de dados PostgreSQL usando o driver "postgres"
	if err != nil {
		panic(err) // panic é usado para interromper imediatamente o software se algo inesperado acontecer
	}

	err = conn.Ping() // testa a conexão com banco de dados, se a conexão for bem sucedida a variavel err será nil, caso contrário ela armazenará o erro que ocorreu 

	return conn, err // retorna a conexão com o banco de dados e qualquer erro que possa ter ocorrido
}
