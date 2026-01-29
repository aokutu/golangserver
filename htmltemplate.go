package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Name string
	Age int
}

func main() {
	tmpl := template.Must(template.ParseFiles("pages/htmltemplate.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data User
		data.Name = "Andrew"
		data.Age =25

		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":8080", nil)
}
