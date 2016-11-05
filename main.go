package main

import (
	"github.com/manawasp/elitekeyboards-watcher/email"
	kbs "github.com/manawasp/elitekeyboards-watcher/keyboards"
)

const SAVE_FILE string = "keyboards.json"
const URL_EK string = "https://elitekeyboards.com/products.php?sub=topre_keyboards,rftenkeyless"

func main() {
	// Get new stats from the website
	keyboards := kbs.WebParse(URL_EK)

	// Load previous stats and compare them
	arr := kbs.Diff(keyboards, kbs.PreviousState(SAVE_FILE))
	if len(arr) > 0 {
		email.Send(arr)
		kbs.Save(SAVE_FILE, keyboards)
	}
}
