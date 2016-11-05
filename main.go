package main

import (
	"github.com/manawasp/elitekeyboards-watcher/email"
	kbs "github.com/manawasp/elitekeyboards-watcher/keyboards"
)

const SENDGRID_KEY string = "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"

func main() {
	DB := "keyboards.json"
	HTML := "email/template.html"
	URL := "https://elitekeyboards.com/products.php?sub=topre_keyboards,rftenkeyless"

	// Get new stats from the website
	keyboards := kbs.WebParse(URL)

	// Load previous stats and compare them
	arr := kbs.Diff(keyboards, kbs.PreviousState(DB))
	if len(arr) > 0 {
		email.Send(SENDGRID_KEY, HTML, arr)
		kbs.Save(DB, keyboards)
	}
}
