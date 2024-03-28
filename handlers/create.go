package handlers // Pacote handlers para agrupar funções de manipulação de solicitações HTTP

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    "github.com/API-POSTGRESQL/models" // Importa o pacote models para interagir com o banco de dados
)

// Create manipula solicitações POST para criar um novo recurso
func Create(w http.ResponseWriter, r *http.Request) {
    var todo models.Todo // Cria uma variável do tipo Todo para armazenar os dados recebidos
    // Decodifica o JSON do corpo da solicitação e o armazena na variável todo
    err := json.NewDecoder(r.Body).Decode(&todo)
    if err != nil {
        // Se houver um erro ao decodificar o JSON, registra o erro e retorna um erro interno do servidor
        log.Printf("Erro ao fazer decode do json: %v", err)
        http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
        return
    }

    // Tenta inserir o todo no banco de dados e recebe o ID da tarefa inserida
    id, err := models.Insert(todo)
    var resp map[string]interface{} // Define um mapa para a resposta JSON

    // Verifica se ocorreu algum erro durante a inserção
    if err != nil {
        resp = map[string]interface{}{ // Se houver um erro, define uma resposta com erro
            "Error":   true,
            "Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err),
        }
    } else {
        resp = map[string]interface{}{ // Se a inserção for bem-sucedida, define uma resposta com sucesso
            "Error":   false,
            "Message": fmt.Sprintf("Todo inserido com sucesso! ID: %d", id),
        }
    }

    // Define o tipo de conteúdo da resposta como JSON
    w.Header().Set("Content-Type", "application/json")
    // Codifica a resposta JSON e a envia de volta ao cliente
    json.NewEncoder(w).Encode(resp)
}