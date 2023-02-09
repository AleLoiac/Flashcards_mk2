package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type flashcard struct {
	Term       string `json:"term"`
	Definition string `json:"definition"`
	Mistakes   int
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

	fmt.Printf("The pair (\"%v\":\"%v\") has been added.\n", f.Term, f.Definition)
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
			fmt.Println("The card has been removed.The card has been removed.")
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

func importCards(reader *bufio.Reader) {

	fmt.Println("File name:")

	fileName, _ := reader.ReadString('\n')
	fileName = strings.TrimSpace(fileName)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("File not found.")
	}
	defer file.Close()

	var importedCards []flashcard
	decoder := json.NewDecoder(file)
	err2 := decoder.Decode(&importedCards)
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	fmt.Println(importedCards)

	for i := len(importedCards) - 1; i >= 0; i-- {
		for j := len(flashcardDeck) - 1; j >= 0; j-- {
			if importedCards[i].Term == flashcardDeck[j].Term {
				flashcardDeck = append(flashcardDeck[:j], flashcardDeck[j+1:]...)
			}
		}
		flashcardDeck = append(flashcardDeck, importedCards[i])
	}

	fmt.Printf("%v cards have been loaded.\n", len(importedCards))

}

func playGame(reader *bufio.Reader) {

	fmt.Println("How many times to ask?")

	num, _ := reader.ReadString('\n')
	num = strings.TrimSpace(num)

	number, _ := strconv.Atoi(num)

	for i := 0; i < number; i++ {

		rand.Seed(time.Now().UnixNano())
		randomInt := rand.Intn(len(flashcardDeck))

		fmt.Printf("Print the definition of \"%v\":\n", flashcardDeck[randomInt].Term)
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)

		if answer == flashcardDeck[randomInt].Definition {
			fmt.Println("Correct!")
		} else {
			flashcardDeck[randomInt].Mistakes++
			control := false
			for j := range flashcardDeck {
				if answer == flashcardDeck[j].Definition {
					fmt.Printf("Wrong. The right answer is \"%v\", but your definition is correct for \"%v\".\n", flashcardDeck[randomInt].Definition, flashcardDeck[j].Term)
					control = true
					break
				}
			}
			if control == false {
				fmt.Printf("Wrong. The right answer is \"%v\".\n", flashcardDeck[randomInt].Definition)
			}
		}
	}

}

func hardest() {

}

func reset() {
	for i := range flashcardDeck {
		flashcardDeck[i].Mistakes = 0
	}
	fmt.Println("Card statistics have been reset.")
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
			importCards(reader)
		case "export":
			exportCards(reader)
		case "ask":
			playGame(reader)
		case "log":

		case "hardest card":
			hardest()
		case "reset stats":
			reset()
		case "exit":
			fmt.Println("Bye bye!")
			fmt.Println(flashcardDeck)
			return
		default:
			fmt.Println("Invalid command, please try again")
		}
	}

}
