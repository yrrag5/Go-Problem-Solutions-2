# Go-Problem-Solutions-2

# Problem sheet 2 solutions 

# Curl Command
 
 curl -i localhost:8080
 
 This command is used in cmder with the curl.exe placed in the bin folder. You must first execute go run of the selected file (Q1.go) before running curl. This will then examine the reponse of the http.ListenandServe.
 
 # Bootstrap 
 
 I based my serve and guess route around the bootstrap stater template: https://getbootstrap.com/docs/4.0/getting-started/introduction/#starter-template
 I included the Url in my index.html to my guess.html which includes a form for user input on selecting a number between 1 and 20.
 
 # Cookies 
 
 I implemented cookies to set the guess as the target for the answer of the guessing game. I insure that its assigned as a integer value and only picks between 1 and 20 using my generateNum variable and comparing it with the user inputted integer. The message "Congratulations!" displays when the user enters the correct guess.   
