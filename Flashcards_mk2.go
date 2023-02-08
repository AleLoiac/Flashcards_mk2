package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type flashcard struct {
	Term       string `json:"term"`
	Definition string `json:"definition"`
}

var flashcardDeck []flashcard

func createCard(reader *bufio.Reader) {

	var f flashcard

	fmt.Println("The card:")
Loop1:
	ter, _ := reader.ReadString('\n')
	ter = strings.TrimSpace(ter)

	for j := range flashcardDeck {
		if ter == flashcardDeck[j].Term {
			fmt.Printf("The term \"%v\" already exists. Try again:\n", flashcardDeck[j].Term)
			goto Loop1
		}
	}

	fmt.Println("The definition of the card:")
Loop2:
	def, _ := reader.ReadString('\n')
	def = strings.TrimSpace(def)

	for z := range flashcardDeck {
		if def == flashcardDeck[z].Definition {
			fmt.Printf("The definition \"%v\" already exists. Try again:\n", flashcardDeck[z].Definition)
			goto Loop2
		}
	}

	f.Term = ter
	f.Definition = def

	flashcardDeck = append(flashcardDeck, f)
}

func removeCard(reader *bufio.Reader) {

	fmt.Println("Which card?")

	term, _ := reader.ReadString('\n')
	term = strings.TrimSpace(term)

	control := false

	for i := len(flashcardDeck) - 1; i >= 0; i-- {
		if term == flashcardDeck[i].Term {
			flashcardDeck = append(flashcardDeck[:i], flashcardDeck[i+1:]...)
			control = true
			break
		}
	}
	if control == false {
		fmt.Printf("Can't remove \"%v\": there is no such card.\n", term)
	}
}

func exportCards(reader *bufio.Reader) {

	fmt.Println("File name:")

	fileName, _ := reader.ReadString('\n')
	fileName = strings.TrimSpace(fileName)

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	FlashcardsJSON, err2 := json.Marshal(flashcardDeck)
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	_, err = file.Write(FlashcardsJSON)
	if err != nil {
		fmt.Println(err)
		return
	}

	var count int

	for range flashcardDeck {
		count++
	}
	fmt.Printf("%v cards have been saved.\n", count)

	// txt export
	//for _, card := range flashcardDeck {
	//	_, err = fmt.Fprintln(file, card.term, card.definition) // writes each card of the 'flashcardDeck' slice
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}

}

func main() {

	flashcardDeck = make([]flashcard, 0)

	for {
		fmt.Println("Input the action (add, remove, import, export, ask, exit):")

		reader := bufio.NewReader(os.Stdin)

		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		switch command {
		case "add":
			createCard(reader)
		case "remove":
			removeCard(reader)
		case "import":

		case "export":
			exportCards(reader)
		case "ask":

		case "exit":
			fmt.Println("Bye bye!")
			fmt.Println(flashcardDeck)
			return
		default:
			fmt.Println("Invalid command, please try again")
		}
	}

}
