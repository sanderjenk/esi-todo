package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Todo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

var todos []Todo

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Accept"})
	myRouter.HandleFunc("/", homePage).Methods("GET", "OPTIONS")
	myRouter.HandleFunc("/todos", getTodos).Methods("GET")
	myRouter.HandleFunc("/todos/{id}", getTodo).Methods("GET")
	myRouter.HandleFunc("/todos", createTodo).Methods("POST")
	myRouter.HandleFunc("/todos/{id}", updateTodo).Methods("PUT")
	myRouter.HandleFunc("/todos/{id}", deleteTodo).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(origins, methods, headers)(myRouter)))
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(todos)
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	found := false
	for _, todo := range todos {
		if todo.Id == key {
			found = true
			json.NewEncoder(w).Encode(todo)
		}
	}

	if !found {
		w.WriteHeader(http.StatusNotFound)
	}
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	json.NewDecoder(r.Body).Decode(&todo)
	todo.Id = strconv.Itoa(len(todos) + 1)
	todos = append(todos, todo)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var newTodo Todo
	json.NewDecoder(r.Body).Decode(&newTodo)
	newTodo.Id = key

	fmt.Println(newTodo)

	var index int
	var found bool
	for i, todo := range todos {
		if todo.Id == key {
			index = i
			found = true
		}
	}

	if found {
		todos = remove(todos, index)
	}

	todos = append(todos, newTodo)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var found bool
	var index int
	for i, todo := range todos {
		if todo.Id == key {
			index = i
			found = true
		}
	}

	if !found {
		w.WriteHeader(http.StatusNotFound)
	} else {
		todos = remove(todos, index)
		w.WriteHeader(http.StatusNoContent)
	}
}

func remove(slice []Todo, s int) []Todo {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	todos = []Todo{
		{Id: "1", Name: "Clean room", Done: false},
		{Id: "2", Name: "Work out", Done: false},
		{Id: "3", Name: "Cook carbonara", Done: false},
	}
	handleRequests()
}
