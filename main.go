package main

import (
	log "github.com/Sirupsen/logrus"

	"github.com/manawasp/elitekeyboards-watcher/email"
	kbs "github.com/manawasp/elitekeyboards-watcher/keyboards"
)

func main() {
	// Get new stats from the website
	keyboards := kbs.WebParse()

	// Load previous stats and compare them
	arr := kbs.Diff(keyboards, kbs.PreviousState())

	log.Println(len(arr))
	if len(arr) > 0 {
		email.Send(arr)
		kbs.Save(keyboards)
	}
}
