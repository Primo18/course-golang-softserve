package main

// @title TODO API
// @version 1.0
// @description API para gestionar TODOs
// @host localhost:8080
// @BasePath /
import (
	"fmt"
	"log"
	"myapp/internal/todos"
	"net/http"

	_ "myapp/docs" // Esto se generará automáticamente

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @Summary Bienvenida
// @Description Muestra la página de bienvenida
// @Tags general
// @Accept json
// @Produce html
// @Success 200 {string} string "Página de bienvenida"
// @Router / [get]
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
        <h1>¡Bienvenido a la API de TODOs! 👋</h1>
        <p>Rutas disponibles:</p>
        <ul>
            <li>GET /todos - Lista todos los TODOs</li>
            <li>GET /todos/{id} - Obtiene un TODO específico</li>
            <li>POST /todos - Crea un nuevo TODO</li>
            <li>PUT /todos/{id} - Actualiza un TODO existente</li>
        </ul>
    `)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Ruta para la documentación Swagger
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	// Ruta principal
	r.Get("/", welcomeHandler)

	// Rutas de TODOs
	r.Post("/todos", todos.CreateTodoHandler)
	r.Get("/todos/{id}", todos.GetTodoHandler)
	r.Get("/todos", todos.ListTodoHandler)
	r.Put("/todos/{id}", todos.UpdateTodoHandler)

	port := ":8080"
	fmt.Printf("🚀 Server running on http://localhost%s\n", port)
	fmt.Printf("📚 Swagger documentation: http://localhost%s/swagger/index.html\n", port)

	log.Fatal(http.ListenAndServe(port, r))
}
