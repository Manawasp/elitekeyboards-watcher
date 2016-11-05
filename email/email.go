package email

import (
	"bytes"
	"text/template"

	log "github.com/Sirupsen/logrus"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	kbs "github.com/manawasp/elitekeyboards-watcher/keyboards"
	"github.com/manawasp/elitekeyboards-watcher/utils"
)

func Send(sendgridKey, pathTemplate string, d []kbs.Keyboard) {
	// generate template
	var buf bytes.Buffer
	t, _ := template.ParseFiles(utils.GetExecDir() + pathTemplate)
	t.Execute(&buf, d)
	s := buf.String()
	// Config message
	from := mail.NewEmail("Elitekeyboards - Notification", "clovis.kyndt@gmail.com")
	subject := "EliteKeyboards Product Availability"
	to := mail.NewEmail("Clovis Kyndt", "clovis.kyndt@gmail.com")
	content := mail.NewContent("text/html", s)
	m := mail.NewV3MailInit(from, subject, to, content)
	// send email
	request := sendgrid.GetRequest(sendgridKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Email sent!")
		log.Println(response.StatusCode)
		log.Println(response.Body)
		log.Println(response.Headers)
	}
}
