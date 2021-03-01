import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map } from 'rxjs/operators';

@Injectable()
export class QuoteService {

  constructor(private http: HttpClient) { }

  getQuote() {
    return this.http.get("https://goquotes-api.herokuapp.com/api/v1/random?count=1").pipe(
      map((result: any) => result.quotes[0]),
    )
  }
}
