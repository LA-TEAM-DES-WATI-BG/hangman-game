package main

import "fmt"

func letter_true() {
	var letter string
	var a []rune
	for {
		fmt.Scan(&letter)
		a = []rune(letter)
		if ((a[0] > 96 && a[0] < 123) || (a[0] > 64 && a[0] < 91)) && len(a) < 2 {

			fmt.Printf("Vous avez choisit %c", a[0])
			break
		} else {
			fmt.Println("choisi une seule lettre de l'alphabet ")
		}
	}
}
