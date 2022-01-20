package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type homePage struct {
	Strng string
}
type lvl struct {
	Lvl1 string
	Lvl2 string
	Lvl3 string
}

type Letter struct {
	Value string
	Used  bool
}

type donnée struct {
	Oui      []string
	Letters  [26]Letter
	Letters2 [5]Letter
}

var templates = template.Must(template.ParseFiles("HangmanHTML/hangman.html"))
var templates2 = template.Must(template.ParseFiles("HangmanHTML/lvl1.html"))
var templates3 = template.Must(template.ParseFiles("HangmanHTML/lvl2.html"))
var templates4 = template.Must(template.ParseFiles("HangmanHTML/lvl3.html"))
var data donnée
var ChoixBot string = rdmWord()
var NbErrror int

func handler(w http.ResponseWriter, r *http.Request) {
	homeP := homePage{
		Strng: "hangman games",
	}
	templates.Execute(w, homeP)
}

func lvl_1(w http.ResponseWriter, r *http.Request) {
	alphabet2 := [5]string{"A", "B", "C", "D", "E"}
	for g := 0; g < 5; g++ {
		ms := Letter{
			Value: alphabet2[g],
			Used:  false,
		}
		data.Letters2[g] = ms
	}
	letter := r.FormValue("letter")
	alphabet := [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	for g := 0; g < len(alphabet); g++ {
		ms := Letter{
			Value: alphabet[g],
			Used:  false,
		}
		if data.Letters[g].Used == true {
			ms = Letter{
				Value: alphabet[g],
				Used:  true,
			}
		}
		data.Letters[g] = ms
	}
	for i := 0; i < len(alphabet); i++ {
		if data.Letters[i].Value == letter {
			data.Letters[i].Used = true
			y := letterTrue(data.Letters[i].Value, ChoixBot)
			if y == false {
				NbErrror += 1
			}
		}
	}
	fmt.Println(NbErrror)
	templates2.Execute(w, data)
}

func letterTrue(Lettre, ChoixBot string) bool {
	ChoixBot2 := []rune(ChoixBot)
	var Mot []string
	for i := 0; i < len(ChoixBot2); i++ {
		var conver string
		conver = string(ChoixBot2[i] - 32)
		Mot = append(Mot, conver)
	}
	fmt.Println(Mot)
	for i := 0; i < len(ChoixBot2); i++ {
		if Lettre == Mot[i] {

			return true
		}
	}
	return false
}

func lvl_2(w http.ResponseWriter, r *http.Request) {
	lvl2 := lvl{
		Lvl2: "test",
	}
	templates3.Execute(w, lvl2)
}

func lvl_3(w http.ResponseWriter, r *http.Request) {
	lvl3 := lvl{
		Lvl3: "test",
	}
	templates4.Execute(w, lvl3)
}
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

func main() {
	fs := http.FileServer(http.Dir("./image"))
	http.Handle("/image/", http.StripPrefix("/image", fs))

	fs2 := http.FileServer(http.Dir("./css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs2))
	// on a parametrer les path pour link les files enssemble
	//
	//
	http.HandleFunc("/", handler)
	http.HandleFunc("/lvl1", lvl_1)
	http.HandleFunc("/lvl2", lvl_2)
	http.HandleFunc("/lvl3", lvl_3)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
