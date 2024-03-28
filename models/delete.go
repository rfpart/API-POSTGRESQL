package models // Pacote models para agrupar estruturas de dados relacionadas

import "github.com/API-POSTGRESQL/db" // Importa o pacote db para interagir com o banco de dados

// Delete exclui uma tarefa do banco de dados com o ID fornecido
func Delete(id int64) (int64, error) {
    // Abre uma conexão com o banco de dados
    conn, err := db.OpenConnection()
    if err != nil {
        return 0, err // Retorna 0 e o erro se não for possível abrir a conexão
    }
    defer conn.Close() // Fecha a conexão quando a função terminar, independentemente do resultado

    // Executa a exclusão no banco de dados com a instrução SQL fornecida
    res, err := conn.Exec(`DELETE FROM todos WHERE id=$1`, id)
    if err != nil {
        return 0, err // Retorna 0 e o erro se ocorrer um erro durante a execução da exclusão
    }

    // Retorna o número de linhas afetadas pela exclusão
    return res.RowsAffected()
}

