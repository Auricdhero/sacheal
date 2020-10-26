package main

import (
    "html/template"
    "net/http"
    "fmt"
)

type ContactDetails struct {
	fName string
	Email   string
	Telephone string
    Subject string
    Message string
    To string
}

func main() {
    tmpl := template.Must(template.ParseFiles("contact.html"))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            tmpl.Execute(w, nil)
            return
        }

        details := ContactDetails{
			fName:	r.FormValue("form-name"),
			Email:   r.FormValue("form-email"),
			Telephone: r.FormValue("form-tel"),
            Subject: r.FormValue("form-subject"),
            Message: r.FormValue("form-message"),
        }

        To:= "datauric@gmail.com"
        Message:= ("Name:"+ fName +"\nEmail:" + Email +"\nTel:" +Telephone +"\nSub:" +Subject + "\nMessage:" + Message)
        
       
        // do something with details
        _ = details

        tmpl.Execute(w, struct{ Success bool }{true})
    })


    http.ListenAndServe(":8080", nil)
}