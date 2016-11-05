package keyboards

import (
	"encoding/json"
	"io/ioutil"

	"github.com/manawasp/elitekeyboards-watcher/utils"
)

func PreviousState(path string) (keyboards Keyboards) {
	file, e := ioutil.ReadFile(utils.GetExecDir() + path)
	if e == nil {
		json.Unmarshal(file, &keyboards)
	}
	return keyboards
}
