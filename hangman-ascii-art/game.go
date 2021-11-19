package main

import (
	"fmt"
	"strconv"
	"time"
)

var (
	lettreU   string
	NbrOfGame = nbrOfGame()
	langueB   = language()
	pseudo    = Pseudo()
	coupF     = 0
)

func Game() {
	for NbrOfGame >= 1 {
		timer := time.Now()
		var Tofind = mot()
		var Found []string
		var life = 10
		n := len(Tofind)/2 - 1
		for i := 0; i < len(Tofind); i++ {
			if string(Tofind[i]) == " " {
				Found = append(Found, " ")
			} else {
				Found = append(Found, "_")
			}
		}
		for j := 0; j < n; {
			r := randInt(len(Tofind))
			letter := Tofind[r]
			for f := range Tofind {
				if Tofind[f] == letter {
					Found[f] = string(letter)
					j++
				}
			}
		}
		end := false
		fmt.Println()
		Repr(Found)
		fmt.Println()
		for end == false {
			word0 := word()
			for i := range word0 {
				if string(word0[i]) < "A" || string(word0[i]) > "Z" && string(word0[i]) < "a" || string(word0[i]) > "z" || len(word0) == 0 || len(word0) > len(Tofind) {
					if langueB == false {
						fmt.Println("Mauvaise entrée")
					} else if langueB == true {
						fmt.Println("Wrong input")
					}
				}
			}
			if try(word0, Tofind) == true {
				if len(word0) == 1 {
					for p := range Tofind {
						if string(Tofind[0]) == ToUpper(word0) {
							Found[0] = ToUpper(word0)
						}
						if string(Tofind[p]) == word0 {
							Found[p] = word0
						}
					}
				} else if len(word0) > 1 {
					for i := range word0 {
						for j := range Tofind {
							if word0[i] == Tofind[j] {
								Found[j] = string(word0[i])
							}
						}
					}
				}

				Repr(Found)
				fmt.Println()
				for _, let := range Found {
					if let == "_" {
						end = false
						break
					} else {
						end = true
					}
				}
				fmt.Print("\n")
			} else if try(word0, Tofind) == false {
				if len(word0) == 1 {
					life--
					coupF++
				} else if life > 1 {
					life -= 2
					coupF += 2
				}
				repr(life)
				fmt.Println()
				Repr(Found)
				fmt.Println()
				if langueB == false {
					fmt.Println()
					fmt.Printf("Pas présent, il vous reste %d essaies", life)
					fmt.Println()
					fmt.Print("Lettres déjà utilisées : ")
					fmt.Print(lettreU)
				} else if langueB == true {
					fmt.Println()
					fmt.Printf("Not present in the word, %d attempts remainin", life)
					fmt.Println()
					fmt.Print("Letters already used : ")
					fmt.Print(lettreU)
				}
				fmt.Println()
				if life == 0 {
					end = true
				}
			}
		}
		if life > 0 {
			if langueB == false {
				fmt.Println("Bien Joué!")
				fmt.Printf("Vous avez fait %d erreurs, ", coupF)
				fmt.Printf("la partie a duré %s", time.Since(timer))
				fmt.Println()
				fmt.Println()
			} else if langueB == true {
				fmt.Println("Good Job!")
				fmt.Printf("You've done %d errors, ", coupF)
				fmt.Printf("the game lasted %s", time.Since(timer))
				fmt.Println()
				fmt.Println()
			}
		} else {
			if langueB == false {
				fmt.Println("Tu es pendu! Dommage :(")
				fmt.Printf("tu devais trouver le mot %s", Tofind)
				fmt.Println()
				fmt.Printf("La partie a duré %s", time.Since(timer))
			} else if langueB == true {
				fmt.Println("You're hang! Sad :(")
				fmt.Printf("The word was %s", Tofind)
				fmt.Println()
				fmt.Printf("The game lasted %s", time.Since(timer))
			}

			fmt.Println()
			fmt.Println()
		}
		score += strconv.Itoa(coupF) + ", "
		score += time.Since(timer).String() + "\n"
		NbrOfGame--
	}

	Score()
}
