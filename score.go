package main

import (
	"os"
)

func Score() {
	//scoreL := strings.Split(string(score), ", ")
	file, err := os.OpenFile("score.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	defer func(file *os.File) {
		_ = file.Close()
	}(file) // on ferme automatiquement à la fin de notre programme
	if err != nil {
		panic(err)
	}
	_, err = file.WriteString(score)
	if err != nil {
		panic(err)
	}
}
