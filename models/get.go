package models

import "github.com/API-POSTGRESQL/db" // Importa o pacote db para interagir com o banco de dados

// Get recupera uma tarefa do banco de dados com base no ID fornecido
func Get(id int64) (todo Todo, err error) {
	// Abre uma conexão com o banco de dados
	conn, err := db.OpenConnection()
	if err != nil {
		return // Retorna um erro se não for possível abrir a conexão
	}
	defer conn.Close() // Fecha a conexão quando o método terminar, independentemente do resultado

	// Define a consulta SQL para selecionar uma tarefa com base no ID
	row := conn.QueryRow(`SELECT * FROM todos WHERE id=$1`, id)

	// Escaneia os valores retornados pela consulta e atribui-os à estrutura todo
	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

	return // Retorna a tarefa recuperada do banco de dados e qualquer erro que ocorra durante a execução da consulta
}
