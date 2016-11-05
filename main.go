package main

import (
	log "github.com/Sirupsen/logrus"
)

func main() {
	// Get new stats from the website
	keyboards := newStats()
	// Load Previous Stats
	k := extractStats()
	// Compare New Stats and Previous
	arr := compareStats(keyboards, k)
	log.Println(len(arr))
	if len(arr) > 0 {
		sendEmail(arr)
		saveStats(keyboards)
	}
}
