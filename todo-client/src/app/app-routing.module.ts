import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { CreateTodoComponent } from './create-todo/create-todo.component';
import { ListTodoComponent } from './list-todo/list-todo.component';

const routes: Routes = [
  {
    path: "todos", component: ListTodoComponent
  },
  {
    path: "todos/new", component: CreateTodoComponent
  },
  {
    path: "", redirectTo: "todos", pathMatch: "full"
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes, { relativeLinkResolution: 'legacy' })],
  exports: [RouterModule]
})
export class AppRoutingModule { }
