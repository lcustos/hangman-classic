package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func language() bool {
	choix := false
	lang := false
	for choix == false {
		fmt.Println("Choose your language : Eng (1) or Fr (2)")
		var language string
		_, _ = fmt.Scan(&language)
		if language == "1" || language == "ENG" || language == "eng" || language == "Eng" {
			lang = true
			choix = true
			score += "English, "
		} else if language == "2" || language == "FR" || language == "Fr" || language == "fr" {
			lang = false
			choix = true
			score += "Français, "
		}
	}
	return lang
}

func Pseudo() string {
	var pseudo string
	verif := false
	for verif == false {
		if langueB == true {
			fmt.Println("What is your gamename ?")
		} else {
			fmt.Println("Quel est votre pseudo ?")
		}
		_, _ = fmt.Scanln(&pseudo)
		if len(pseudo) != 0 {
			score += pseudo + ", "
			verif = true
		}
	}
	return pseudo
}

var hangman = GetHangman()

func repr(life int) {
	colorReset := "\033[0m"
	colorRed := "\033[31m"
	fmt.Println(colorRed, hangman[life+1], colorReset)
}

func GetHangman() []string {
	var hangman []string
	var Tsplit []string
	content, err := ioutil.ReadFile("ressources/hangman.txt")
	if err != nil {
		log.Fatal(err)
	} else {
		Tsplit = strings.Split(string(content), "\n\n")
		for _, i := range Tsplit {
			hangman = append(hangman, i)
		}
		for i, j := 0, len(hangman)-1; i < j; i, j = i+1, j-1 {
			hangman[i], hangman[j] = hangman[j], hangman[i]
		}
	}
	return hangman
}

func nbrOfGame() int {
	var nbr int
	if langueB == true {
		fmt.Println("How many games would you play ?")
	} else if langueB == false {
		fmt.Println("Combien de partie voulez-vous jouez ?")
	}
	_, _ = fmt.Scanln(&nbr)
	return nbr
}

func randInt(max int) int {
	min := 0
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

var score string

func Score() {
	file, err := os.OpenFile("score.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	if err != nil {
		panic(err)
	}
	_, err = file.WriteString(score)
	if err != nil {
		panic(err)
	}
}
