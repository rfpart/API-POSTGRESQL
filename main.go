package main

import (
	"fmt"
	"net/http"

	"github.com/API-POSTGRESQL/handlers"         // Importa os manipuladores de solicitação HTTP
	"github.com/aprendagolang/api-pgsql/configs" // Importa as configurações do servidor
	"github.com/go-chi/chi/v5"                   // Importa o roteador de solicitações HTTP
)

func main() {
	// Carrega as configurações do servidor
	err := configs.Load()
	if err != nil {
		panic(err) // Em caso de erro ao carregar as configurações, encerra o programa
	}

	// Cria um novo roteador Chi
	r := chi.NewRouter()

	// Define as rotas e seus manipuladores correspondentes
	r.Post("/", handlers.Create)       // Rota para criar um novo recurso
	r.Put("/{id}", handlers.Update())  // Rota para atualizar um recurso existente
	r.Delete("/{id}", handlers.Delete) // Rota para excluir um recurso existente
	r.Get("/", handlers.List)          // Rota para obter todos os recursos
	r.Get("/{id}", handlers.Get)       // Rota para obter um recurso específico

	// Inicia o servidor na porta especificada nas configurações
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
