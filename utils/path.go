package utils

import (
	"os"
	"strings"
)

func GetExecDir() (rez string) {
	arr := strings.Split(os.Args[0], "/")
	i := 0
	for i < len(arr)-1 {
		rez += arr[i] + "/"
		i += 1
	}
	return rez
}
