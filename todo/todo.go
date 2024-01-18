package todo

import (
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Todo struct {
	Slug string `json:"slug"`
	Title string `json:"title"`
	Body string `json:"body"`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", GetAllTodos)
	router.Post("/", CreateTodo)
	router.Get("/{todoId}", GetTodo)
	router.Delete("/{todoId}", DeleteTodo)
	return router
}

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos := []*Todo{
		&Todo{"1", "Todo 1", "Todo 1 body"},
		&Todo{"2", "Todo 2", "Todo 2 body"},
	}
	render.JSON(w, r, todos) // A chi router helper for serializing and returning json
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Todo created successfully"
	render.JSON(w, r, response)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Todo  deleted successfully"
	render.JSON(w, r, response)
}

func GetTodo (w http.ResponseWriter, r *http.Request) {
	todoId := chi.URLParam(r, "todoId")
	response := make(map[string]string)
	response["message"] = "Todo with id " + todoId + " fetched successfully"
	render.JSON(w, r, response)
}