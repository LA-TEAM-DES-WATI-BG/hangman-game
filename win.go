package main

func win(nblettrerestante []int, a string) bool {
	tab := []rune(a)
	if len(nblettrerestante) == len(tab) {
		return false
	}
	return true
}
