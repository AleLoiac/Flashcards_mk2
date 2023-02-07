package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type flashcard struct {
	term       string
	definition string
}

var flashcardDeck []flashcard

func createCard(reader *bufio.Reader) {

	var f flashcard

	fmt.Println("The card:")
Loop1:
	ter, _ := reader.ReadString('\n')
	ter = strings.TrimSpace(ter)

	for j := range flashcardDeck {
		if ter == flashcardDeck[j].term {
			fmt.Printf("The term \"%v\" already exists. Try again:\n", flashcardDeck[j].term)
			goto Loop1
		}
	}

	fmt.Println("The definition of the card:")
Loop2:
	def, _ := reader.ReadString('\n')
	def = strings.TrimSpace(def)

	for z := range flashcardDeck {
		if def == flashcardDeck[z].definition {
			fmt.Printf("The definition \"%v\" already exists. Try again:\n", flashcardDeck[z].definition)
			goto Loop2
		}
	}

	f.term = ter
	f.definition = def

	flashcardDeck = append(flashcardDeck, f)
}

func removeCard(reader *bufio.Reader) {

	fmt.Println("Which card?")

	term, _ := reader.ReadString('\n')
	term = strings.TrimSpace(term)

	control := false

	for i := len(flashcardDeck) - 1; i >= 0; i-- {
		if term == flashcardDeck[i].term {
			flashcardDeck = append(flashcardDeck[:i], flashcardDeck[i+1:]...)
			control = true
			break
		}
	}
	if control == false {
		fmt.Printf("Can't remove \"%v\": there is no such card.\n", term)
	}
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
