package main

import "fmt"

func affiche(tab []int, mot string) {

	lettre := []rune(mot)
	var secret []rune
	for i := 0; i < len(lettre); i++ {
		secret = append(secret, '_')
	}
	for i := 0; i < len(tab); i++ {
		secret[tab[i]] = lettre[tab[i]]
	}
	for i := 0; i < len(lettre); i++ {
		fmt.Printf("%c ", secret[i])
	}
}
