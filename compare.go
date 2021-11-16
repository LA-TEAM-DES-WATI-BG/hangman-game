package main

import "fmt"

func compare(user_choice rune, bot_choice string) []int {
	bot_choice_rune := []rune(bot_choice)
	var tab []int
	if user_choice >= 65 && user_choice <= 90 {
		user_choice = user_choice + 32
	}
	for i := 0; i < len(bot_choice_rune); i++ {
		if bot_choice_rune[i] == user_choice {
			tab = append(tab, i)
		}
	}
	fmt.Println(tab)
	return tab
}
