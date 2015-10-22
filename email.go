package main

import (
    "bytes"
    "text/template"
    "github.com/sendgrid/sendgrid-go"
    "fmt"
)

func sendEmail(d []Keyboard) {
    // generate template
    var buf bytes.Buffer
    t, _ := template.ParseFiles("email_alert_beautify.html")
    t.Execute(&buf, d)
    s := buf.String()
    // send email
    sg := sendgrid.NewSendGridClientWithApiKey("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")
    message := sendgrid.NewMail()
    message.AddTo("clovss.mna@gmail.com")
    message.AddToName("Clovis Kyndt")
    message.SetSubject("EliteKeyboards Product Availability")
    message.SetHTML(s)
    message.SetFrom("clovss.mna@gmail.com")
    if r := sg.Send(message); r == nil {
        fmt.Println("Email sent!")
    } else {
        fmt.Println(r)
    }
}
