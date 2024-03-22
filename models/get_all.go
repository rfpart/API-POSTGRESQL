package models // Pacote models para agrupar estruturas de dados relacionadas

import "github.com/API-POSTGRESQL/db" // Importa o pacote db para interagir com o banco de dados

// GetAll retorna todas as tarefas do banco de dados
func GetAll() (todos []Todo, err error) {
	// Abre uma conexão com o banco de dados
	conn, err := db.OpenConnection()
	if err != nil {
		return // Retorna um erro se não for possível abrir a conexão
	}
	defer conn.Close() // Fecha a conexão quando a função terminar, independentemente do resultado

	// Executa a consulta SQL para selecionar todas as tarefas
	rows, err := conn.Query(`SELECT * FROM todos`)
	if err != nil {
		return // Retorna um erro se não for possível executar a consulta
	}
	defer rows.Close() // Fecha o cursor quando a função terminar, independentemente do resultado

	// Loop para escanear os valores das colunas de cada linha
	for rows.Next() {
		var todo Todo

		// Escaneia os valores retornados pela consulta e atribui-os à estrutura todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue // Pula para a próxima iteração do loop se ocorrer um erro ao escanear os valores
		}

		todos = append(todos, todo) // Adiciona a tarefa à lista de tarefas recuperadas
	}

	return // Retorna todas as tarefas recuperadas e qualquer erro que ocorra durante a execução da consulta
}
