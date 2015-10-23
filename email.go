package main

import (
  "bytes"
  "text/template"
  "github.com/sendgrid/sendgrid-go"
  "fmt"
)

const NOTIFY_TPL string =  "email_alert_beautify.html"
const SENDGRID_KEY string = "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"

func sendEmail(d []Keyboard) {
    // generate template
    var buf bytes.Buffer
    fmt.Println(getExecDir())
    fmt.Println(d)
    t, _ := template.ParseFiles(getExecDir() + NOTIFY_TPL)
    t.Execute(&buf, d)
    s := buf.String()
    // send email
    sg := sendgrid.NewSendGridClientWithApiKey(SENDGRID_KEY)
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
