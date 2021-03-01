import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable()
export class TodoService {
  constructor(private http: HttpClient) { }

  getTodos() {
    return this.http.get("http://localhost:8000/todos");
  }
  getTodo(id: string) {
    return this.http.get(`http://localhost:8000/todos/${id}`);
  }
  createTodo(todo: any) {
    return this.http.post(`http://localhost:8000/todos`, todo);
  }
  updateTodo(id: string, todo: any) {
    return this.http.put(`http://localhost:8000/todos/${id}`, todo);
  }
  deleteTodo(id: string) {
    return this.http.delete(`http://localhost:8000/todos/${id}`);
  }
  completeTodo(todo: any) {
    todo.done = true;
    return this.updateTodo(todo.id, todo);
  }
}