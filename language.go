package main

import (
	"fmt"
	"strings"
)

var (
	langueB = language()
)

func language() bool {
	choix := false
	lang := false
	for choix == false {
		fmt.Println("Choose your language : Eng or Fr")
		var language string
		_, _ = fmt.Scan(&language)
		if strings.ToUpper(language) == "ENG" {
			lang = true
			choix = true
			score += "English, "
		} else if strings.ToUpper(language) == "FR" {
			lang = false
			choix = true
			score += "Français, "
		}
	}
	return lang
}
