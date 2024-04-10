package handlers // Pacote handlers para agrupar funções de manipulação de solicitações HTTP

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/API-POSTGRESQL/models" // Importa o pacote models para interagir com o banco de dados
)

// Get é uma função que manipula solicitações HTTP GET para obter um recurso
func Get(w http.ResponseWriter, r *http.Request) {
	// Extrai o ID do parâmetro da URL e converte para um inteiro
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		// Se houver um erro ao converter o ID, registra o erro e retorna um erro interno do servidor
		log.Printf("Erro ao fazer parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Chama a função Get do pacote models para obter o registro do banco de dados
	todo, err := models.Get(int64(id))
	if err != nil {
		// Se houver um erro ao obter o registro, registra o erro e retorna um erro interno do servidor
		log.Printf("Erro ao obter registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Define o tipo de conteúdo da resposta como JSON
	w.Header().Set("Content-Type", "application/json")

	// Codifica o registro em JSON e envia como resposta HTTP
	json.NewEncoder(w).Encode(todo)
}
