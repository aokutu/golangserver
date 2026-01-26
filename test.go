package main

import (
    "fmt"
    "net/http"
	"os"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    fmt.Fprintln(w, "Welcome to the Home page")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "This is the About page")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "This is the Contact page")
}

func checkname(w http.ResponseWriter, r *http.Request) {

    // Allow only POST requests
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    } 
 
    // Parse the form
    err := r.ParseForm()
    if err != nil {
        http.Error(w, "Error parsing form", http.StatusBadRequest)
        return
    }

	type details struct{
		username string
		password string
	}

	var user details 

	user.username = r.FormValue("name")
	user.password = r.FormValue("password")

	err = os.WriteFile("/jsondata/users.json", []byte("Hello Go"), 0644)
	if err != nil {
	fmt.Fprintln(w, "FAILED TO RECORD", user.username)
	} 
  


	 

    // Print it in the browser
    fmt.Fprintln(w, "You entered:", user)
}


func administrator(w http.ResponseWriter, r *http.Request) {
     http.ServeFile(w, r, "pages/admin.html")
}


func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/about", aboutHandler)
    http.HandleFunc("/contact", contactHandler)
	 http.HandleFunc("/checkname", checkname)
	http.HandleFunc("/administrator", administrator)

    fmt.Println("Server started at http://localhost:8090")
    http.ListenAndServe(":8090", nil)
}
