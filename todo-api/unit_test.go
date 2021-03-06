// Example based on Rest and unit testing in https://golangdocs.com/

package main

import (
	todo "esi-todo/todo-api/todo"
	"testing"
)

func TestAddNewTodo(t *testing.T) {
	testTodo := todo.Todo{
		Id:   "142",
		Name: "Do homework",
		Done: false,
	}
	addNewTodo(testTodo)
	foundTodo, found := findTodo("142")
	if !found || foundTodo != testTodo {
		t.Error("Couldn't find ToDo after adding.")
		return
	}
}

func TestChangeTodo(t *testing.T) {
	testTodo := todo.Todo{
		Id:   "142",
		Name: "Do homework",
		Done: false,
	}
	addNewTodo(testTodo)
	modifiedTodo := todo.Todo{
		Id:   "142",
		Name: "Go jogging",
		Done: false,
	}
	changeTodo("142", modifiedTodo)
	foundTodo, found := findTodo("142")
	if !found || foundTodo.Name == "Go jogging" {
		t.Error("ToDo changes are not implemented.")
		return
	}
}

func TestRemoveTodo(t *testing.T) {
	testTodo := todo.Todo{
		Id:   "142",
		Name: "Do homework",
		Done: false,
	}
	addNewTodo(testTodo)
	removeTodo("142")
	foundTodo, found := findTodo("142")
	if !found || foundTodo == testTodo {
		t.Error("Removing ToDo failed.")
		return
	}
}
