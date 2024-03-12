package models // Pacote models para agrupar estruturas de dados relacionadas

// Todo representa uma tarefa a ser concluída
type Todo struct {
	ID          int64  `json:"id"`          // ID único da tarefa
	Title       string `json:"title"`       // Título da tarefa
	Description string `json:"description"` // Descrição da tarefa
	Done        bool   `json:"done"`        // Indica se a tarefa foi concluída (true) ou não (false)
}
