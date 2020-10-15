package main

import (
    "html/template"
    "net/http"
)

type ContactDetails struct {
	fName string
	Email   string
	Telephone string
    Subject string
    Message string
}

func main() {
    tmpl := template.Must(template.ParseFiles("contact.html"))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            tmpl.Execute(w, nil)
            return
        }

        details := ContactDetails{
			fName:	r.FormValue("fName"),
			Email:   r.FormValue("email"),
			Telephone: r.FormValue("telephone"),
            Subject: r.FormValue("subject"),
            Message: r.FormValue("message"),
        }

        // do something with details
        _ = details

        tmpl.Execute(w, struct{ Success bool }{true})
    })

    http.ListenAndServe(":8080", nil)
}