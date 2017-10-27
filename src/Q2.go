// Author: Garry Cummins
// Date: 23/10/2017


package main

import (
  "net/http"
	"fmt"
)

func main() {
  http.HandleFunc("/", guess)
  http.ListenAndServe(":8080", nil)
}

func guess(w http.ResponseWriter, r *http.Request) {
	//Adding header tags to guessing game
	w.Header().Set("Server", "Guessing game")
  	w.WriteHeader(200)
	//Adding header tags to guessing game
  	fmt.Fprint(w, "<h1>Guessing Game</h1>")
}