package keyboards

func Diff(source, cible Keyboards) (arr []Keyboard) {
	for key, _ := range source {
		_, exist := cible[key]
		if !exist || (cible[key].Available != source[key].Available) {
			arr = append(arr, source[key])
		}
	}
	return
}
