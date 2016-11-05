package keyboards

import (
	"github.com/BurntSushi/toml"
	log "github.com/Sirupsen/logrus"
)

// Parse the "DB" file to retrieve previous state of the website
func PreviousState(path string) *State {
	state := &State{}
	if _, err := toml.DecodeFile("keyboards.toml", state); err != nil {
		log.Errorf("Error: Unable to decode config file, %v", err)
	}
	// file, e := ioutil.ReadFile(utils.GetExecDir() + path)
	// if e == nil {
	// 	json.Unmarshal(file, &keyboards)
	// }
	return state
}
