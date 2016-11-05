package keyboards

import (
	"encoding/json"
	"io/ioutil"

	"github.com/manawasp/elitekeyboards-watcher/utils"
)

func Save(keyboards Keyboards) {
	b, _ := json.Marshal(keyboards)
	err := ioutil.WriteFile(utils.GetExecDir()+SAVE_FILE, b, 0644)
	if err != nil {
		panic(err)
	}
}
