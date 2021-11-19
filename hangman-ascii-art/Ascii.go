package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func GetAscii() []string {
	var ascii []string
	var Tsplit []string
	content, err := ioutil.ReadFile("hangman-ascii-art/ressources/standard.txt")
	if err != nil {
		log.Fatal(err)
	} else {
		Tsplit = strings.Split(string(content), "\r\n\r\n")
		for _, i := range Tsplit {
			ascii = append(ascii, i)
		}
	}
	return ascii
}

func Map() map[string]string {
	mapp := map[string]string{}
	for i := 32; i <= 126; i++ {
		mapp[string(rune(i))] = GetAscii()[i-32]
	}

	return mapp
}

func Repr(i []string) {
	for k := 0; k < 8; k++ {
		for j := 0; j < len(i); j++ {
			Tsplit2 := strings.Split(Map()[i[j]], "\r\n")
			fmt.Print(Tsplit2[k])

		}
		fmt.Println()
	}
}
