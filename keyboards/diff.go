package keyboards

// Diff compare 2 Keyboards availability and return a new array of keyboards if
// the availability change
func Diff(source, cible *State) (arr []Keyboard) {
	for key, _ := range source.Keyboards {
		_, exist := cible.Keyboards[key]
		if !exist || (cible.Keyboards[key].Available != source.Keyboards[key].Available) {
			arr = append(arr, source.Keyboards[key])
		}
	}
	return
}
