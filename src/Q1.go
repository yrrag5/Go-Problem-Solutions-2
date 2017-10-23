// Author: Garry Cummins
// Date: 23/10/2017

//Adapted from : http://www.alexedwards.net/blog/golang-response-snippets

package main

//Imports with http
import (
  "net/http"
  "fmt"
)

func main() {
  http.HandleFunc("/", guess)
  // Curl to run port 8080 
  http.ListenAndServe(":8080", nil)
}

func guess(w http.ResponseWriter, r *http.Request) {
  //Setting server name 
  w.Header().Set("Server", "Guessing game")
  w.WriteHeader(200)
  // Responds with text guessing game 
  fmt.Fprint(w, "Guessing Game")
}