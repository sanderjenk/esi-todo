import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { TodoService } from '../todo.service';

@Component({
  selector: 'app-create-todo',
  templateUrl: './create-todo.component.html',
  styleUrls: ['./create-todo.component.css'],
  providers: [TodoService]
})
export class CreateTodoComponent implements OnInit {
  todo: any;

  todoForm = new FormGroup({
    name: new FormControl('', Validators.required),
    done: new FormControl(false)
  });
  constructor(
    private todoService: TodoService, 
    private router: Router, 
    private route: ActivatedRoute) { }

  ngOnInit(): void {
  }

  onSubmit() {
    this.todoService.createTodo(this.todoForm.value).subscribe(() => {
      this.router.navigate(['../'], {relativeTo: this.route})
    })
  }

}
