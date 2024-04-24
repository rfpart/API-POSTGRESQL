package handlers // Pacote handlers para agrupar funções de manipulação de solicitações HTTP

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/API-POSTGRESQL/models" // Importa o pacote models para interagir com o banco de dados
	"github.com/go-chi/chi/v5"
)

// Update manipula solicitações PUT para atualizar um recurso
func Update(w http.ResponseWriter, r *http.Request) {
	// Extrai o ID do parâmetro da URL e converte para um inteiro
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		// Se houver um erro ao converter o ID, registra o erro e retorna um erro interno do servidor
		log.Printf("Erro ao fazer parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Cria uma variável do tipo Todo para armazenar os dados atualizados
	var todo models.Todo

	// Decodifica o JSON do corpo da solicitação e o armazena na variável todo
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		// Se houver um erro ao decodificar o JSON, registra o erro e retorna um erro interno do servidor
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Chama a função Update do pacote models para atualizar o registro no banco de dados
	rows, err := models.Update(int64(id), todo)
	if err != nil {
		// Se houver um erro ao atualizar o registro, registra o erro e retorna um erro interno do servidor
		log.Printf("Erro ao atualizar registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Verifica se mais de um registro foi atualizado e registra um aviso se necessário
	if rows > 1 {
		log.Printf("Error: foram atualizados %d registros", rows)
	}

	// Cria uma resposta JSON indicando que os dados foram atualizados com sucesso
	resp := map[string]interface{}{
		"Error":   false,
		"Message": "dados atualizados com sucesso!",
	}

	// Define o tipo de conteúdo da resposta como JSON e envia a resposta JSON de volta ao cliente
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
