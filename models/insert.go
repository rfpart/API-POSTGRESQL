package models // Pacote models para agrupar estruturas de dados relacionadas

import (
	"github.com/API-POSTGRESQL/db" // Importa o pacote db para interagir com o banco de dados
)

// Insert insere uma nova tarefa no banco de dados
func Insert(todo Todo) (id int64, err error) {
	// Abre uma conexão com o banco de dados
	conn, err := db.OpenConnection()
	if err != nil {
		return // Retorna um erro se não for possível abrir a conexão
	}
	defer conn.Close() // Fecha a conexão quando a função terminar, independentemente do resultado

	// Define a instrução SQL para inserir uma nova tarefa na tabela 'todos'
	sql := `INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURN id`

	// Executa a instrução SQL, passando os valores correspondentes e escaneando o ID retornado
	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)

	return // Retorna o ID da nova tarefa e qualquer erro que ocorra durante a execução da consulta
}
