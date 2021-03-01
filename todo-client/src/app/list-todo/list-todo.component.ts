import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { QuoteModalComponent } from '../quote-modal/quote-modal.component';
import { QuoteService } from '../quote.service';
import { TodoService } from '../todo.service';
import { concatMap } from 'rxjs/operators';

@Component({
  selector: 'app-list-todo',
  templateUrl: './list-todo.component.html',
  styleUrls: ['./list-todo.component.css'],
  providers: [TodoService, QuoteService]
})
export class ListTodoComponent implements OnInit {
  todos = []

  constructor(
    private todoService: TodoService,
    private router: Router, 
    private route: ActivatedRoute,
    private modalService: NgbModal,
    private quoteService: QuoteService
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
    this.todoService.completeTodo({...todo}).pipe(
      concatMap(() => this.quoteService.getQuote()),
      concatMap((quote: any) => {
        const modalRef = this.modalService.open(QuoteModalComponent);
        modalRef.componentInstance.quote = quote;
        return modalRef.result;
      })
    ).subscribe(() => this.getTodos())
  }

  add() {
    this.router.navigate(['new'], {relativeTo: this.route})
  }
}
