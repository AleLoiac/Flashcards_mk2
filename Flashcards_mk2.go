package main

import "fmt"

type flashcard struct {
	term       string
	definition string
}

var flashcardDeck []flashcard

func createCard() {

}

func main() {

	for {
		fmt.Println("Input the action (add, remove, import, export, ask, exit):")

		var command string
		_, err := fmt.Scan(&command)
		if err != nil {
			return
		}

		switch command {
		case "add":
			createCard()
		case "remove":

		case "import":

		case "export":

		case "ask":

		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid command, please try again")
		}
	}

}
