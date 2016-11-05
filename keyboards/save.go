package keyboards

import (
	"encoding/json"
	"io/ioutil"

	"github.com/manawasp/elitekeyboards-watcher/utils"
)

// Save rewrite the "DB" file with the current state of keyboards
func Save(path string, keyboards Keyboards) {
	b, _ := json.Marshal(keyboards)
	err := ioutil.WriteFile(utils.GetExecDir()+path, b, 0644)
	if err != nil {
		panic(err)
	}
}
