package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math"
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

var programLog []string

func createCard(reader *bufio.Reader) {

	var f flashcard

	l3 := fmt.Sprintln("The card:")
	fmt.Print(l3)
	programLog = append(programLog, l3)
Loop1:
	ter, _ := reader.ReadString('\n')
	ter = strings.TrimSpace(ter)
	programLog = append(programLog, ter)

	for j := range flashcardDeck {
		if ter == flashcardDeck[j].Term {
			l4 := fmt.Sprintf("The term \"%v\" already exists. Try again:\n", flashcardDeck[j].Term)
			fmt.Print(l4)
			programLog = append(programLog, l4)
			goto Loop1
		}
	}

	l5 := fmt.Sprintln("The definition of the card:")
	fmt.Print(l5)
	programLog = append(programLog, l5)
Loop2:
	def, _ := reader.ReadString('\n')
	def = strings.TrimSpace(def)
	programLog = append(programLog, def)

	for z := range flashcardDeck {
		if def == flashcardDeck[z].Definition {
			l6 := fmt.Sprintf("The definition \"%v\" already exists. Try again:\n", flashcardDeck[z].Definition)
			fmt.Print(l6)
			programLog = append(programLog, l6)
			goto Loop2
		}
	}

	f.Term = ter
	f.Definition = def

	flashcardDeck = append(flashcardDeck, f)

	l7 := fmt.Sprintf("The pair (\"%v\":\"%v\") has been added.\n", f.Term, f.Definition)
	fmt.Print(l7)
	programLog = append(programLog, l7)
}

func removeCard(reader *bufio.Reader) {

	l8 := fmt.Sprintln("Which card?")
	fmt.Print(l8)
	programLog = append(programLog, l8)

	term, _ := reader.ReadString('\n')
	term = strings.TrimSpace(term)
	programLog = append(programLog, term)

	control := false

	for i := len(flashcardDeck) - 1; i >= 0; i-- {
		if term == flashcardDeck[i].Term {
			flashcardDeck = append(flashcardDeck[:i], flashcardDeck[i+1:]...)
			control = true
			l9 := fmt.Sprintln("The card has been removed.The card has been removed.")
			fmt.Print(l9)
			programLog = append(programLog, l9)
			break
		}
	}
	if control == false {
		l10 := fmt.Sprintf("Can't remove \"%v\": there is no such card.\n", term)
		fmt.Print(l10)
		programLog = append(programLog, l10)
	}
}

func exportCards(reader *bufio.Reader) {

	l11 := fmt.Sprintln("File name:")
	fmt.Print(l11)
	programLog = append(programLog, l11)

	fileName, _ := reader.ReadString('\n')
	fileName = strings.TrimSpace(fileName)
	programLog = append(programLog, fileName)

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
	l12 := fmt.Sprintf("%v cards have been saved.\n", count)
	fmt.Print(l12)
	programLog = append(programLog, l12)

	// txt export
	//for _, card := range flashcardDeck {
	//	_, err = fmt.Fprintln(file, card.term, card.definition) // writes each card of the 'flashcardDeck' slice
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}

}

func importCards(reader *bufio.Reader) {

	l13 := fmt.Sprintln("File name:")
	fmt.Print(l13)
	programLog = append(programLog, l13)

	fileName, _ := reader.ReadString('\n')
	fileName = strings.TrimSpace(fileName)
	programLog = append(programLog, fileName)

	file, err := os.Open(fileName)
	if err != nil {
		l14 := fmt.Sprintln("File not found.")
		fmt.Print(l14)
		programLog = append(programLog, l14)
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

	l15 := fmt.Sprintf("%v cards have been loaded.\n", len(importedCards))
	fmt.Print(l15)
	programLog = append(programLog, l15)

}

func playGame(reader *bufio.Reader) {

	l16 := fmt.Sprintln("How many times to ask?")
	fmt.Print(l16)
	programLog = append(programLog, l16)

	num, _ := reader.ReadString('\n')
	num = strings.TrimSpace(num)
	programLog = append(programLog, num)

	number, _ := strconv.Atoi(num)

	for i := 0; i < number; i++ {

		rand.Seed(time.Now().UnixNano())
		randomInt := rand.Intn(len(flashcardDeck))

		l17 := fmt.Sprintf("Print the definition of \"%v\":\n", flashcardDeck[randomInt].Term)
		fmt.Print(l17)
		programLog = append(programLog, l17)
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)
		programLog = append(programLog, answer)

		if answer == flashcardDeck[randomInt].Definition {
			l18 := fmt.Sprintln("Correct!")
			fmt.Print(l18)
			programLog = append(programLog, l18)
		} else {
			flashcardDeck[randomInt].Mistakes++
			control := false
			for j := range flashcardDeck {
				if answer == flashcardDeck[j].Definition {
					l19 := fmt.Sprintf("Wrong. The right answer is \"%v\", but your definition is correct for \"%v\".\n", flashcardDeck[randomInt].Definition, flashcardDeck[j].Term)
					fmt.Print(l19)
					programLog = append(programLog, l19)
					control = true
					break
				}
			}
			if control == false {
				l20 := fmt.Sprintf("Wrong. The right answer is \"%v\".\n", flashcardDeck[randomInt].Definition)
				fmt.Print(l20)
				programLog = append(programLog, l20)
			}
		}
	}

}

func hardest() {

	maxErrors := 0
	for i := 0; i <= len(flashcardDeck)-1; i++ {
		maxErrors = int(math.Max(float64(flashcardDeck[i].Mistakes), float64(maxErrors)))
	}
	//fmt.Println(maxErrors)

	var hardestCards []flashcard
	hardestCards = make([]flashcard, 0)
	for i := range flashcardDeck {
		if flashcardDeck[i].Mistakes == maxErrors {
			hardestCards = append(hardestCards, flashcardDeck[i])
		}
	}
	if maxErrors == 0 {
		l21 := fmt.Sprintln("There are no cards with errors.")
		fmt.Print(l21)
		programLog = append(programLog, l21)
	} else if len(hardestCards) == 1 {
		l22 := fmt.Sprintf("The hardest card is \"%v\". You have %v errors answering it\n", hardestCards[0].Term, maxErrors)
		fmt.Print(l22)
		programLog = append(programLog, l22)
	} else {
		l23 := fmt.Sprintf("The hardest cards are ")
		fmt.Print(l23)
		programLog = append(programLog, l23)
		for i := range hardestCards {
			l24 := fmt.Sprintf("\"%v\" ", hardestCards[i].Term)
			fmt.Print(l24)
			programLog = append(programLog, l24)
		}
		l25 := fmt.Sprintf("You have %v errors answering them.\n", maxErrors)
		fmt.Print(l25)
		programLog = append(programLog, l25)
	}
}

func reset() {
	for i := range flashcardDeck {
		flashcardDeck[i].Mistakes = 0
	}
	l26 := fmt.Sprintln("Card statistics have been reset.")
	fmt.Print(l26)
	programLog = append(programLog, l26)
}

func saveLog(reader *bufio.Reader) {

	l27 := fmt.Sprintln("File name:")
	fmt.Print(l27)
	programLog = append(programLog, l27)

	fileName, _ := reader.ReadString('\n')
	fileName = strings.TrimSpace(fileName)
	programLog = append(programLog, fileName)

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for _, v := range programLog {
		_, err = fmt.Fprintln(file, v)
		if err != nil {
			log.Fatal(err)
		}
	}

	l28 := fmt.Sprintln("The log has been saved.")
	fmt.Print(l28)
	programLog = append(programLog, l28)
}

func main() {

	flashcardDeck = make([]flashcard, 0)

	for {
		l := fmt.Sprintln("Input the action (add, remove, import, export, ask, exit, log, hardest card, reset stats):")
		fmt.Print(l)
		programLog = append(programLog, l)

		reader := bufio.NewReader(os.Stdin)

		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)
		programLog = append(programLog, command)

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
			saveLog(reader)
		case "hardest card":
			hardest()
		case "reset stats":
			reset()
		case "exit":
			l1 := fmt.Sprintln("Bye bye!")
			fmt.Print(l1)
			programLog = append(programLog, l1)

			fmt.Println(flashcardDeck)
			return
		default:
			l2 := fmt.Sprintln("Invalid command, please try again")
			fmt.Print(l2)
			programLog = append(programLog, l2)
		}
	}

}
