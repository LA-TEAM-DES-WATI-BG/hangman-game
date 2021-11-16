package main

import "fmt"

func main() {
	var user_choice rune
	var tab []int
	var position []int
	var wina bool
	bot_choice := rdmWord()
	for nbError := 10; nbError > 0; {
		hangmanStatus(nbError)
		fmt.Println()
		user_choice = letterTrue()
		position = compare(user_choice, bot_choice)
		for i := 0; i < len(position); i++ {
			tab = append(tab, position[i])
		}
		if len(position) == 0 {

			nbError--
		}
		affiche(tab, bot_choice)
		wina = win(tab, bot_choice)
		if wina == false {
			fmt.Println("vous avez win")
			break
		}
		// 	func erreur()
		// 	func  reussite()
	}
}
