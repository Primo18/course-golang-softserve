package todos

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// @Summary Crear TODO
// @Description Crea un nuevo TODO
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body Todo true "Datos del TODO"
// @Success 201 {object} Todo "TODO creado"
// @Failure 400 {string} string "Error en la solicitud"
// @Router /todos [post]
func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	todo.ID = len(todos) + 1
	todos = append(todos, todo)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)

}

// @Summary Obtener TODO
// @Description Obtiene un TODO por su ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "ID del TODO"
// @Success 200 {object} Todo "TODO encontrado"
// @Failure 400 {string} string "ID inválido"
// @Failure 404 {string} string "TODO no encontrado"
// @Router /todos/{id} [get]
func GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for _, todo := range todos {
		if todo.ID == id {
			json.NewEncoder(w).Encode(todo)
			return
		}
	}

	http.NotFound(w, r)

}

// @Summary Listar TODOs
// @Description Obtiene la lista de todos los TODOs
// @Tags todos
// @Accept json
// @Produce json
// @Success 200 {array} Todo "Lista de TODOs"
// @Router /todos [get]
func ListTodoHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(todos)
}

// @Summary Actualizar TODO
// @Description Actualiza un TODO existente por su ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "ID del TODO"
// @Param todo body Todo true "Datos actualizados del TODO"
// @Success 200 {object} Todo "TODO actualizado"
// @Failure 400 {string} string "Error en la solicitud"
// @Failure 404 {string} string "TODO no encontrado"
// @Router /todos/{id} [put]
func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i, todo := range todos {
		if todo.ID == id {
			var updatedTodo Todo
			if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			todos[i].Completed = updatedTodo.Completed
			json.NewEncoder(w).Encode(todos[i])
			return
		}
	}
	http.NotFound(w, r)

}
