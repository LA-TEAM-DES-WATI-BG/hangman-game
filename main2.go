package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
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
	Oui     []string
	Letters []Letter
}

var templates = template.Must(template.ParseFiles("HangmanHTML/hangman.html"))
var templates2 = template.Must(template.ParseFiles("HangmanHTML/lvl1.html"))
var templates3 = template.Must(template.ParseFiles("HangmanHTML/lvl2.html"))
var templates4 = template.Must(template.ParseFiles("HangmanHTML/lvl3.html"))
var donnéeVar donnée

func handler(w http.ResponseWriter, r *http.Request) {
	homeP := homePage{
		Strng: "hangman games",
	}
	templates.Execute(w, homeP)
}

func lvl_1(w http.ResponseWriter, r *http.Request) {

	letter := r.FormValue("letter")
	// alphabet := [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	// for i, fdp := range data.Letters {
	// 	if fdp.Value == letter {
	// 		data.Letters[i] = Letter{Value: fdp.Value, Used: true}
	// 		break
	// 	}
	data := donnée{
		Letters: []Letter{
			{Value: "A", Used: false},
		},
	}
	for i, fdp := range donnéeVar.Letters {
		if fdp.Value == letter {
			donnéeVar.Letters[i] = Letter{
				Value: fdp.Value,
				Used:  true,
			}
		}
	}
	fmt.Println(letter)

	templates2.Execute(w, data)
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
