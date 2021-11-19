package main

import "fmt"

var (
	NbrOfGame = nbrOfGame()
)

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
