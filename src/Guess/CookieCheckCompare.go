// Author: Garry Cummins
// Date: 27/10/2017

//Adapted from: https://astaxie.gitbooks.io/build-web-application-with-golang/en/06.1.html

package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type text struct {
	message string
	guess string
}

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "guess.html")
}

// Checking for a cookie and comparing it with guess 
func guessGameHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm();
	g := r.FormValue("guess")

	rand.Seed(int64(time.Now().Nanosecond()))

	if _, err := r.Cookie("target"); err != nil {
		// Generating random number between 1 and 20 
		var randNum = ((rand.Int() % 19) + 1)

		i, _ := strconv.Atoi(g)

		// Checking if guess matches the number selected 
		if i == randNum{
			g = "Correct!"
		}// Inner if 

		num := strconv.Itoa(randNum)

		// Declaring cookie and setting the name as target
		cookie := http.Cookie{
			Name:  "target",
			Value: num,
		}// Cookie 

		http.SetCookie(w, &cookie)
	}// if

	m := text{message: "Guess a number between 1 and 20: ", guess: g}
	t, _ := template.ParseFiles("guess.tmpl")
	t.Execute(w, m)
}// Func

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/guess", guessGameHandler)
	http.ListenAndServe(":8080", nil)
}