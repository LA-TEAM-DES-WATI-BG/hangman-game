package main

import "fmt"

func main1() {
	//var copie []int
	var user_choice rune
	var tab []int
	var position []int
	var wina bool
	var tabletter []rune
	bot_choice := rdmWord()
	for nbError := 10; nbError > -10; {
		hangmanStatus(nbError)
		fmt.Println()
		user_choice = letterTrue(tabletter)
		position = compare(user_choice, bot_choice)
		tabletter = append(tabletter, user_choice)
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
		fmt.Println(tab)
		if nbError <= 0 {
			fmt.Printf("Vous avez perdu le véritable mot était %s \n", bot_choice)
			break
		}

	}
}
