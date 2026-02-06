package main

import (
	"log"
	"net/http"
	"fmt"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pages/ajaxtest.html")
	})

	

	http.HandleFunc("/processdata", func(w http.ResponseWriter, r *http.Request) {

		username:= r.FormValue("username")
		fmt.Fprintln(w, username)
	})


	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
