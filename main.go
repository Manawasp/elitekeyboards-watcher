package main

import (
	"os"

	"github.com/BurntSushi/toml"
	log "github.com/Sirupsen/logrus"

	"github.com/manawasp/elitekeyboards-watcher/email"
	kbs "github.com/manawasp/elitekeyboards-watcher/keyboards"
)

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

	// Get api sendgrid key
	if len(conf.SendgridKey) > 0 {
		conf.SendgridKey = os.Getenv("SENDGRID_API_KEY")
	}

	// Get new stats from the website
	keyboards, err := kbs.WebParse(conf.URL)
	if err != nil {
		log.Errorf("Error: kbs.WebParse, %v.", err)
	}

	// Load previous stats and compare them
	arr := kbs.Diff(keyboards, kbs.PreviousState(conf.DB))
	if len(arr) > 0 {
		email.Send(conf.SendgridKey, conf.HTML, arr)
		kbs.Save(conf.DB, keyboards)
	}
}
