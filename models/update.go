package models // Pacote models para agrupar estruturas de dados relacionadas

import "github.com/API-POSTGRESQL/db" // Importa o pacote db para interagir com o banco de dados

// Update atualiza uma tarefa existente no banco de dados com o ID fornecido
func Update(id int64, todo Todo) (int64, error) {
    // Abre uma conexão com o banco de dados
    conn, err := db.OpenConnection()
    if err != nil {
        return 0, err // Retorna 0 e o erro se não for possível abrir a conexão
    }
    defer conn.Close() // Fecha a conexão quando a função terminar, independentemente do resultado

    // Executa a atualização no banco de dados com a instrução SQL fornecida
    res, err := conn.Exec(`UPDATE todos SET title=$2, description=$3, done=$4 WHERE id=$1`, id, todo.Title, todo.Description, todo.Done)
    if err != nil {
        return 0, err // Retorna 0 e o erro se ocorrer um erro durante a execução da atualização
    }

    // Retorna o número de linhas afetadas pela atualização
    return res.RowsAffected()
}
