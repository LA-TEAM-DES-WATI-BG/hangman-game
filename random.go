package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

func rdmWord() string {
	rand.Seed(time.Now().UTC().UnixNano())
	data, err := ioutil.ReadFile("words3.txt")
	if err != nil {
		fmt.Println(err)
	}
	dat := string(data)
	dat2 := SplitWhiteSpaces(dat)
	Rdmwrd := randInt(0, len(dat2))
	fmt.Printf("le bot a choisit sont mot !")
	return (dat2[Rdmwrd])
}
func SplitWhiteSpaces(args string) []string {
	var i int
	c := []rune(args)
	var mot string
	var slice []string
	min := 0
	for i = 0; i < len(args); i++ {
		if c[min] == 10 {
			min = min + 1
		}
		mot = string(c[min:i])
		if c[min+1] == 10 {
			mot = string(c[min : min+1])
			slice = append(slice, mot)
			min = min + 1
		}
		if i == len(args)-1 {
			mot = string(c[min : i+1])
		}
		if c[i] == 10 || c[i] == 13 {
			slice = append(slice, mot)
			min = i + 1
			i = i + 2
		} else if i == len(args)-1 {
			slice = append(slice, mot)
		}
	}
	return slice
}
func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return (rand.Intn(max))
} //cette fonction permet de choisir un nombre aleatoire.
