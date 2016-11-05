package keyboards

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/manawasp/elitekeyboards-watcher/utils"
)

// Save rewrite the "DB" file with the current state of keyboards
func Save(path string, state *State) error {
	var file, err = os.OpenFile(utils.GetExecDir()+path, os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
		return err
	}
	defer file.Close()

	encoder := toml.NewEncoder(file)
	encoder.Encode(state)
	// b, _ := json.Marshal(keyboards)
	// err := ioutil.WriteFile(, b, 0644)
	return nil
}
