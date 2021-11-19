package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var (
	hangman = GetHangman()
)

func GetHangman() []string {
	var hangman []string
	var Tsplit []string
	content, err := ioutil.ReadFile("ressources/hangman.txt")
	if err != nil {
		log.Fatal(err)
	} else {
		Tsplit = strings.Split(string(content), "\r\n\r\n")
		for _, i := range Tsplit {
			hangman = append(hangman, i)
		}
		for i, j := 0, len(hangman)-1; i < j; i, j = i+1, j-1 {
			hangman[i], hangman[j] = hangman[j], hangman[i]
		}
	}
	return hangman
}

func repr(life int) {
	colorReset := "\033[0m"
	colorRed := "\033[31m"
	fmt.Println(colorRed, hangman[life+1], colorReset)
}
