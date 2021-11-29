package main

import "fmt"

func letterTrue(rep []rune) rune {
	var letter string
	var a []rune
	var erreur int
	for {
		erreur = 0
		fmt.Println("veuillez choisir une lettre")
		if len(rep) > 0 {
			fmt.Printf("Lettres utilisés: ")
			for i := 0; i < len(rep); i++ {
				fmt.Printf("%c ", rep[i])
			}
			fmt.Println("")
		}
		fmt.Scan(&letter)
		a = []rune(letter)
		if ((a[0] > 96 && a[0] < 123) || (a[0] > 64 && a[0] < 91)) && len(a) < 2 {
			fmt.Printf("Vous avez choisit %c\n", a[0])
			for i := 0; i < len(rep); i++ {
				if rep[i] == a[0] {
					erreur = 1
				}
			}
			if erreur == 1 {
				fmt.Println("Choisit une lettre que tu n'a déjà pas choisit")
			} else {
				break
			}
		} else {
			fmt.Println("choisi une seule lettre de l'alphabet ")
		}
	}
	return a[0]
}
