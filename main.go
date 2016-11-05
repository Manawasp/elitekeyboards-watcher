package main

import (
	"github.com/manawasp/elitekeyboards-watcher/email"
	kbs "github.com/manawasp/elitekeyboards-watcher/keyboards"
)

const SAVE_FILE string = "keyboards.json"
const URL_EK string = "https://elitekeyboards.com/products.php?sub=topre_keyboards,rftenkeyless"
const NOTIFY_TPL string = "email/template.html"
const SENDGRID_KEY string = "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"

func main() {
	// Get new stats from the website
	keyboards := kbs.WebParse(URL_EK)

	// Load previous stats and compare them
	arr := kbs.Diff(keyboards, kbs.PreviousState(SAVE_FILE))
	if len(arr) > 0 {
		email.Send(SENDGRID_KEY, NOTIFY_TPL, arr)
		kbs.Save(SAVE_FILE, keyboards)
	}
}
