import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { TodoService } from '../todo.service';

@Component({
  selector: 'app-list-todo',
  templateUrl: './list-todo.component.html',
  styleUrls: ['./list-todo.component.css']
})
export class ListTodoComponent implements OnInit {
  todos = []

  constructor(
    private todoService: TodoService,
    private router: Router, 
    private route: ActivatedRoute
    ) { }

  ngOnInit() {
    this.getTodos();
  }

  getTodos() {
    this.todoService.getTodos().subscribe((items: []) => {
      this.todos = items;
    })
  }

  delete(todo) {
    this.todoService.deleteTodo(todo.id).subscribe(() => {
      this.getTodos()
    });
  }

  complete(todo) {
    this.todoService.completeTodo({...todo}).subscribe(() => {
      this.getTodos()
    });
  }

  add() {
    this.router.navigate(['new'], {relativeTo: this.route})
  }
}
