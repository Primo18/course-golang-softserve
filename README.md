# MyProject: Calculadora y API de Tareas Pendientes

Este proyecto contiene dos funcionalidades principales desarrolladas en Go:

1. **Calculadora**: Contiene una función que realiza una división segura entre dos números, además de pruebas unitarias, benchmarks y fuzz testing.
2. **API de Tareas Pendientes**: Una API REST sencilla que permite gestionar una lista de tareas pendientes (todos) utilizando el framework Chi.

## Estructura del Proyecto

```
COURSE-GOLANG-SOFTSERVE/
├── go.mod
├── cmd/
│   ├── main/
│   │   └── main.go            # Entrada principal para ejecutar el proyecto
│   ├── api/
│   │   └── main.go            # Entrada principal para ejecutar la API
├── internal/
│   ├── calculator/
│   │   ├── divide.go          # Lógica del cálculo
│   │   ├── divide_test.go     # Pruebas unitarias para el cálculo
│   │   └── fuzz_test.go       # Fuzz testing para la función Divide
│   ├── todos/
│   │   ├── todo.go            # Definición del modelo Todo y lógica de operaciones
│   │   ├── todo_handler.go    # Controladores para gestionar los endpoints REST
│   │   ├── todo_test.go       # Pruebas unitarias para la API de tareas
├── pkg/
│   └── errors/
│       └── error_handler.go   # Manejo de errores reutilizable
├── benchmarks/
│   ├── divide_benchmark_test.go # Benchmark para la función Divide
    └── todo_benchmark_test.go   # Benchmark para las operaciones de la API
```

## Ejercicio 1: Calculadora (Divide)

### Enunciado

1. **Crear una función** `Divide(a, b float64) (float64, error)` que devuelva un error si `b` es 0 utilizando el nuevo manejo de errores con `errors.New`.
2. **Escribir pruebas unitarias** para verificar una división válida y el manejo de la división por cero.
3. **Agregar un benchmark** para evaluar el rendimiento de la función.
4. **Implementar fuzz testing** para detectar casos inesperados.

### Cómo Ejecutar

1. **Ejecutar Pruebas Unitarias**:
   ```sh
   go test ./internal/calculator
   ```

2. **Ejecutar Benchmark**:
   ```sh
   go test -bench=. ./benchmarks/divide_benchmark_test.go
   ```

3. **Ejecutar Fuzz Testing**:
   ```sh
   go test -fuzz=Fuzz ./internal/calculator
   ```

## Ejercicio 2: API REST de Tareas Pendientes

### Enunciado

1. **Crear Tarea**: Endpoint `POST /todos` que recibe un objeto Todo en formato JSON con los campos `ID`, `Title`, y `Completed` (booleano) y lo almacena en una lista.
2. **Obtener Tarea**: Endpoint `GET /todos/{id}` que devuelve los detalles de la tarea con el `ID` proporcionado.
3. **Listar Tareas**: Endpoint `GET /todos` que devuelve todas las tareas pendientes.
4. **Actualizar Tarea**: Endpoint `PUT /todos/{id}` que actualiza el estado de `Completed` de una tarea.

### Cómo Ejecutar la API

1. **Instalar Dependencias** (si es necesario):
   ```sh
   go mod tidy
   ```

2. **Levantar el Servidor**:
   ```sh
   go run ./cmd/api/main.go
   ```
   Esto iniciará la API en el puerto 8080.

3. **Probar los Endpoints**:
   Puedes utilizar herramientas como [Postman](https://www.postman.com/) o `curl` para probar los endpoints de la API:
   - Crear una tarea:
     ```sh
     curl -X POST http://localhost:8080/todos -H "Content-Type: application/json" -d '{"title": "Aprender Go", "completed": false}'
     ```
   - Listar tareas:
     ```sh
     curl http://localhost:8080/todos
     ```
   - Obtener tarea por ID:
     ```sh
     curl http://localhost:8080/todos/1
     ```
   - Actualizar el estado de una tarea:
     ```sh
     curl -X PUT http://localhost:8080/todos/1 -H "Content-Type: application/json" -d '{"completed": true}'
     ```

## Requisitos Previos

- **Go 1.18 o superior**: Este proyecto utiliza Go 1.18 o versiones posteriores para aprovechar las funcionalidades de fuzz testing y manejar los módulos de manera adecuada.
- **Dependencias**: El proyecto utiliza `github.com/go-chi/chi/v5` para manejar los endpoints de la API REST.

## Instalación

1. Clona este repositorio:
   ```sh
   git clone <url-del-repositorio>
   cd myproject
   ```
2. Descarga las dependencias:
   ```sh
   go mod tidy
   ```

## Tecnologías Utilizadas

- **Go**: Lenguaje principal para implementar la lógica del cálculo y la API REST.
- **Chi**: Framework HTTP ligero para manejar los endpoints de la API.
- **Go Testing**: Utilizado para escribir pruebas unitarias, benchmarks y fuzz testing.

## Ejecución Completa del Proyecto

Para probar ambas funcionalidades, ejecuta `main.go` para la calculadora y `api/main.go` para la API REST:

- **Calculadora**: Ejecuta el archivo `cmd/main/main.go`:
  ```sh
  go run ./cmd/main/main.go
  ```

- **API de Tareas**: Ejecuta el archivo `cmd/api/main.go`:
  ```sh
  go run ./cmd/api/main.go
  ```

