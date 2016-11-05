package keyboards

import (
	"encoding/json"
	"io/ioutil"

	"github.com/manawasp/elitekeyboards-watcher/utils"
)

func PreviousState() (keyboards Keyboards) {
	file, e := ioutil.ReadFile(utils.GetExecDir() + SAVE_FILE)
	if e == nil {
		json.Unmarshal(file, &keyboards)
	}
	return keyboards
}
