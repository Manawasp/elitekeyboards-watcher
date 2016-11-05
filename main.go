package main

import (
	"github.com/BurntSushi/toml"
	log "github.com/Sirupsen/logrus"

	kbs "github.com/manawasp/elitekeyboards-watcher/keyboards"
)

const SENDGRID_KEY string = "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"

type AppConfig struct {
	DB          string `toml:"DB"`
	HTML        string `toml:"HTML"`
	URL         string `toml:"URL"`
	SendgridKey string `toml:"SENDGRID_API_KEY"`
}

func main() {
	// Retrieve AppConfig
	var conf AppConfig
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		log.Errorf("Error: Unable to decode config file, %v", err)
		return
	}

	// Get new stats from the website
	keyboards, err := kbs.WebParse(conf.URL)
	if err != nil {
		log.Errorf("Error: kbs.WebParse, %v.", err)
	}

	// Load previous stats and compare them
	arr := kbs.Diff(keyboards, kbs.PreviousState(conf.DB))
	log.Println(arr)
	if len(arr) > 0 {
		// email.Send(SENDGRID_KEY, conf.HTML, arr)
		kbs.Save(conf.DB, keyboards)
	}
}
