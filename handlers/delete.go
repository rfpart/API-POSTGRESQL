package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/API-POSTGRESQL/models"
)

// Delete manipula solicitações DELETE para remover um recurso
func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Chama a função Delete do pacote models para remover o registro do banco de dados
	rows, err := models.Delete(int64(id))
	if err != nil {
		log.Printf("Erro ao remover registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Verifica se mais de um registro foi removido e registra um aviso se necessário
	if rows > 1 {
		log.Printf("Error: foram removidos %d registros", rows)
	}

	// Cria uma resposta JSON indicando que o registro foi removido com sucesso
	resp := map[string]interface{}{
		"Error":   false,
		"Message": "registro removido com sucesso!",
	}

	// Define o tipo de conteúdo da resposta como JSON e envia a resposta JSON de volta ao cliente
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
