package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
)


func main()  {
	//setting up gmail credentials
	GMAIL_USERNAME := os.Getenv("datauric@gmail.com")
	GMAIL_PASSWORD := os.Getenv("darlco25")
	gmailAuth := smtp.PlainAuth("", GMAIL_USERNAME, GMAIL_PASSWORD, "smtp.gmail.com")

	//parsefile returns templates and errors
	t, _ := template.ParseFiles("contact.html")

	//body holds email data

	var body bytes.Buffer

	headers := "MINE-VERSION: 1.0; \nContent-Type: text/html;"
	body.Write([]byte(fmt.Sprintf("Subject: CLIENT APPOINTMENT\n%s\n\n", headers)))

	t.Execute(&body, struct {
		fName string
		eMail string
		Tel string
		Subject string
		Message string
	}{
		fName: "Contact Name",
		eMail: "Contact Email",
		Tel: "Contact Telephone",
		Subject: "Contact Subject",
		Message: "Contact Message",
	})
	smtp.SendMail("smpt.gmail.com:587", gmailAuth, GMAIL_USERNAME, []string{"datauric@gmail.com"}, body.Bytes())

}