package test

import (
	"bytes"
	"encoding/json"
	todo "esi-todo/todo-api/todo"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestCreateTodo(t *testing.T) {
	testTodo := todo.Todo{
		Name: "Do laundry",
		Done: false,
	}
	todoJSON, _ := json.Marshal(testTodo)
	resp, err := http.Post("http://localhost:8000/todos", "", bytes.NewBuffer(todoJSON))
	if err != nil {
		t.Error("Problem adding new todo via REST:", err)
		return
	}

	createdTodoJSON, _ := ioutil.ReadAll(resp.Body)
	createdTodo := todo.Todo{}
	json.Unmarshal(createdTodoJSON, &createdTodo)

	resp, err = http.Get("http://localhost:8000/todos/" + createdTodo.Id)

	if err != nil {
		t.Error("Problem reading todo via REST.")
		return
	}
	findTodoJSON, _ := ioutil.ReadAll(resp.Body)
	findTodo := todo.Todo{}
	json.Unmarshal(findTodoJSON, &findTodo)

	if (findTodo.Name != testTodo.Name) || (findTodo.Done != testTodo.Done) {
		t.Error("Couldn't find or parse todo after adding via REST.")
		return
	}
}

func TestDeleteTodo(t *testing.T) {
	testTodo := todo.Todo{
		Name: "Do laundry",
		Done: false,
	}
	todoJSON, _ := json.Marshal(testTodo)
	resp, err := http.Post("http://localhost:8000/todos", "", bytes.NewBuffer(todoJSON))
	if err != nil {
		t.Error("Problem adding new todo via REST:", err)
		return
	}

	createdTodoJSON, _ := ioutil.ReadAll(resp.Body)
	createdTodo := todo.Todo{}
	json.Unmarshal(createdTodoJSON, &createdTodo)

	req, err := http.NewRequest(http.MethodDelete, "http://localhost:8000/todos/"+createdTodo.Id, nil)

	client := &http.Client{}
	resp, err = client.Do(req)

	if err != nil {
		t.Error("Problem deleting todo via REST:", err)
		return
	}

	resp, err = http.Get("http://localhost:8000/todos")
	todosJSON, _ := ioutil.ReadAll(resp.Body)
	todos := []todo.Todo{}
	json.Unmarshal(todosJSON, &todos)

	for _, item := range todos {
		if item.Id == createdTodo.Id {
			t.Error("Problem deleting todo via REST:", err)
			return
		}
	}

}

func TestUpdateTodo(t *testing.T) {
	testTodo := todo.Todo{
		Name: "Do laundry",
		Done: false,
	}
	todoJSON, _ := json.Marshal(testTodo)
	resp, err := http.Post("http://localhost:8000/todos", "", bytes.NewBuffer(todoJSON))
	if err != nil {
		t.Error("Problem adding new todo via REST:", err)
		return
	}

	createdTodoJSON, _ := ioutil.ReadAll(resp.Body)
	createdTodo := todo.Todo{}
	json.Unmarshal(createdTodoJSON, &createdTodo)

	createdTodo.Done = true
	newTodoJSON, _ := json.Marshal(createdTodo)

	req, err := http.NewRequest(http.MethodPut, "http://localhost:8000/todos/"+createdTodo.Id, bytes.NewBuffer(newTodoJSON))

	client := &http.Client{}
	resp, err = client.Do(req)

	if err != nil {
		t.Error("Problem updating todo via REST:", err)
		return
	}

	resp, err = http.Get("http://localhost:8000/todos/" + createdTodo.Id)
	updatedTodoJSON, _ := ioutil.ReadAll(resp.Body)
	updatedTodo := todo.Todo{}
	json.Unmarshal(updatedTodoJSON, &updatedTodo)

	fmt.Println(updatedTodo.Done != true)
	if updatedTodo.Done != true {
		t.Error("Problem getting todo via REST:", err)
		return
	}

}
