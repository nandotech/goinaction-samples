//webforms.go
package main

import (
	"html/template"
	"net/http"
)

//ContactDetails struct form data type
type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {
	tmpl := template.Must(template.ParseFiles("forms.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		details := ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		//Perform action, save details to database, pass to function
		//to send Email, notification, etc.
		_ = details

		tmpl.Execute(w, struct{ Success bool }{true})
	})
	http.ListenAndServe(":8080", nil)
}
