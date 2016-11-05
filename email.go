package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const NOTIFY_TPL string = "email_alert_beautify.html"
const SENDGRID_KEY string = "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"

func sendEmail(d []Keyboard) {
	// generate template
	var buf bytes.Buffer
	t, _ := template.ParseFiles(getExecDir() + NOTIFY_TPL)
	t.Execute(&buf, d)
	s := buf.String()
	// Config message
	from := mail.NewEmail("Elitekeyboards - Notification", "clovis.kyndt@gmail.com")
	subject := "EliteKeyboards Product Availability"
	to := mail.NewEmail("Clovis Kyndt", "clovis.kyndt@gmail.com")
	content := mail.NewContent("text/html", s)
	m := mail.NewV3MailInit(from, subject, to, content)
	// send email
	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Email sent!")
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
