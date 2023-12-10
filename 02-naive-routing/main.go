package main

import (
  "log"
  "net/http"
)

func main() {
  log.Println("Listening on port 3000...")
  log.Fatal(http.ListenAndServe(":3000", http.HandlerFunc(HandleRequest)))
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
  switch r.URL.Path {
  case "/":
    w.Write([]byte(`
      <h1>Home Page</h1>
      <a href="/contact">Go Contact Page</a>
    `))
  case "/contact":
    w.Write([]byte(`
      <h1>Contact Page</h1>
      <a href="/">Go Home Page</a>
    `))
  default:
    // ...404
    w.Write([]byte(`
      Oops! 404
    `))
  }
}
