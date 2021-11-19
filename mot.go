package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const url = "https://random-word-api.herokuapp.com/word?number=50"

type words []string

func mot() string {
	var word string
	difficulty := false
	l := false
	for l == false {
		if langueB == true {
			l = true
			listMot := *new(words).getWords()
			var easy []string
			var medium []string
			var hard []string
			for i := 0; i < len(listMot); i++ {

				if len(listMot[i]) > 1 && len(listMot[i]) <= 5 {
					easy = append(easy, listMot[i])

				} else if len(listMot[i]) > 5 && len(listMot[i]) <= 8 {
					medium = append(medium, listMot[i])

				} else if len(listMot[i]) > 8 && len(listMot[i]) <= 13 {
					hard = append(hard, listMot[i])
				}
			}
			langueB = true
			for difficulty == false {
				fmt.Println("Choose your level: Easy (1), Medium (2), Hard (3)")
				var level string
				_, _ = fmt.Scan(&level)
				if level == "Easy" || level == "easy" || level == "EASY" || level == "1" {
					word = easy[randInt(len(easy))]
					score += "easy, "
					difficulty = true
				} else if level == "Medium" || level == "medium" || level == "MEDIUM" || level == "2" {
					word = medium[randInt(len(medium))]
					score += "medium, "
					difficulty = true
				} else if level == "Hard" || level == "hard" || level == "HARD" || level == "3" {
					word = hard[randInt(len(hard))]
					score += "hard, "
					difficulty = true
				} else {
					fmt.Println("STAY FOCUS !!! ")
				}
			}
		} else if langueB == false {
			l = true
			f := false
			var listMot []string
			for f == false {
				fmt.Println("Choissisez votre fichier : transports.txt (1), animaux.txt (2), nourriture.txt (3)")
				var fichier string
				_, _ = fmt.Scan(&fichier)
				if fichier == "1" {
					fichier = "transports.txt"
				}
				if fichier == "2" {
					fichier = "animaux.txt"
				}
				if fichier == "3" {
					fichier = "nourriture.txt"
				}
				file, err := os.Open(fmt.Sprintf("ressources/%s", fichier))
				if err != nil {
					fmt.Printf("Error when opening file: %s\n", err)
				}
				fileScanner := bufio.NewScanner(file)
				for fileScanner.Scan() {
					f = true
					listMot = append(listMot, fileScanner.Text())
				}
				if err := fileScanner.Err(); err != nil {
					fmt.Printf("Error while reading file: %s\n", err)
				}
				_ = file.Close()
			}
			var facile []string
			var moyen []string
			var difficile []string
			for i := 0; i < len(listMot); i++ {

				if len(listMot[i]) > 1 && len(listMot[i]) <= 5 {
					facile = append(facile, listMot[i])

				} else if len(listMot[i]) > 5 && len(listMot[i]) <= 8 {
					moyen = append(moyen, listMot[i])

				} else if len(listMot[i]) > 8 && len(listMot[i]) <= 13 {
					difficile = append(difficile, listMot[i])
				}
			}
			difficulte := false
			for difficulte == false {
				fmt.Println("Choissisez votre niveau : Facile (1), Moyen (2), Difficile (3)")
				var niveau string
				_, _ = fmt.Scan(&niveau)
				if niveau == "Facile" || niveau == "facile" || niveau == "FACILE" || niveau == "1" {
					word = facile[randInt(len(facile))]
					score += "facile, "
					difficulte = true
				} else if niveau == "Moyen" || niveau == "moyen" || niveau == "MOYEN" || niveau == "2" {
					word = moyen[randInt(len(moyen))]
					score += "moyen, "
					difficulte = true
				} else if niveau == "Difficile" || niveau == "difficile" || niveau == "DIFFICILE" || niveau == "3" {
					word = difficile[randInt(len(difficile))]
					score += "difficile, "
					difficulte = true
				} else {
					fmt.Println("CONCENTRE TOI !!! ")
				}
			}
		}
	}
	score += word + ", "
	return word
}

func word() string {
	first := ""
	if langueB == false {
		fmt.Println("Quel est votre lettre ?")
		_, _ = fmt.Scanln(&first)
	} else if langueB == true {
		fmt.Println("What is your proposition ?")
		_, _ = fmt.Scanln(&first)
	}
	IsPresent := false
	for i := range lettreU {
		if string(lettreU[i]) == first {
			IsPresent = true
		}
	}
	if IsPresent == false {
		if len(lettreU) == 0 {
			lettreU += first
		} else {
			lettreU += ", " + first
		}
	}
	return first
}

func (myList *words) getWords() *words {
	html, _ := http.Get(url)
	res, _ := ioutil.ReadAll(html.Body)
	_ = json.Unmarshal(res, &myList)
	return myList
}

func randInt(max int) int {
	min := 0
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
