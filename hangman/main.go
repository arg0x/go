package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

var randomWord string
var guessedLetters string
var correctLetters []string
var correctGuesses []string
var wrongGuesses []string

var board [7]string
var remLives int
var remStr string

var reader = bufio.NewReader(os.Stdin)

func main() {

	fmt.Println("Welcome to Game")
	fmt.Println("Generated word :", randomWordGen())
	chooseBoard()
	for true {
		showBoard()
		guess := getUserLetter()

		if remLives <= 0 {
			fmt.Println("Sorry, You are Dead!, The secrect word was", randomWord)
			break
		}

		if strings.Contains(randomWord, guess) {
			updateCorrectLetter(guess)
			remStr = fmt.Sprint(correctLetters)
			if strings.Contains(remStr, "") {
				correctGuesses = append(correctGuesses, guess)
				guessedLetters += guess
				if strings.Compare(randomWord, strings.Join(correctLetters, "")) == 0 {
					fmt.Println("\nYou Win! The correct word is ", randomWord)
					break
				}
				fmt.Println("\nCorrect! Now guess the Remaining letters")
			}

		} else {
			wrongGuesses = append(wrongGuesses, guess)
			guessedLetters += guess
			fmt.Println("\nIncorrect!, Try again")

		}

	}

}

func updateCorrectLetter(letter string) {
	runeRandomWord := []rune(randomWord)

	for i, v := range runeRandomWord {
		if string(v) == letter {
			correctLetters[i] = letter
		}
	}
}

func randomWordGen() string {
	seed := time.Now().Unix()
	rand.Seed(seed)
	randomWord = secretWords[rand.Intn(13)]
	correctLetters = make([]string, len(randomWord))
	return randomWord

}

func chooseBoard() {

	for true {
		fmt.Printf("\nPick the way of execution\n1.Gallows\n2.Gulliotine\n3.Quit\n")
		choice, err := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		if err != nil {
			log.Fatal(err)
		} else if choice == "1" {
			fmt.Println("\nChosen Gallows")
			board = gallows
			break

		} else if choice == "2" {
			fmt.Println("\nChosen Gulliotine")
			board = gulliotine
			break

		} else if choice == "3" {
			fmt.Println("\nGame Stopped!")
			os.Exit(0)
		} else {
			fmt.Println("\nEnter a valid choice")
		}
	}

}

func showBoard() {
	fmt.Println()
	fmt.Println(board[len(wrongGuesses)])

	fmt.Printf("\nSecret word : ")

	for _, v := range correctLetters {
		if v == "" {
			fmt.Print("_")
		} else {
			fmt.Print(v)
		}
	}

	fmt.Printf("\nCorrect Guesses : ")

	for _, v := range correctGuesses {
		if v != "" {
			fmt.Print(v + " ")
		}
	}

	fmt.Printf("\nIncorrect Guesses : ")

	for _, v := range wrongGuesses {
		fmt.Print(v + " ")
	}

	remLives = 6 - len(wrongGuesses)

	fmt.Println("\nLives Remaining : ", remLives)

}

func getUserLetter() string {
	var userGuess string
	for true {
		fmt.Println("Guess a Letter : ")
		userGuess, err := reader.ReadString('\n')

		if err == nil {
			userGuess = strings.TrimSpace(userGuess)
			var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
			if len(userGuess) > 1 {
				fmt.Println("Please enter only one letter")
			} else if IsLetter(userGuess) {
				if strings.Contains(guessedLetters, userGuess) {
					fmt.Println("Please enter a letter that you haven't guessed")
				} else {
					return strings.ToUpper(userGuess)
				}
			} else {
				fmt.Println("Please enter a letter")
			}

		} else {
			log.Fatal(err)
		}

	}

	return strings.ToUpper(userGuess)

}

var gallows = [7]string{
	" +------+\n" +
		"        |\n" +
		"        |\n" +
		"        |\n" +
		"        |\n" +
		"        |\n" +
		"        |\n" +
		"        |\n" +
		"/=======|  ",

	" +------+\n" +
		" |      |\n" +
		" |      |\n" +
		"(o)     |\n" +
		"        |\n" +
		"        |\n" +
		"        |\n" +
		"        |\n" +
		"/=======|  ",
	" +------+\n" +
		" |      |\n" +
		" |      |\n" +
		"(O)     |\n" +
		" |      |\n" +
		"        |\n" +
		"        |\n" +
		"        |\n" +
		"/=======|  ",
	" +------+\n" +
		" |      |\n" +
		" |      |\n" +
		"(O)     |\n" +
		"/|      |\n" +
		"        |\n" +
		"        |\n" +
		"        |\n" +
		"/=======|  ",
	" +------+\n" +
		" |      |\n" +
		" |      |\n" +
		"(O)     |\n" +
		"/|\\     |\n" +
		"        |\n" +
		"        |\n" +
		"        |\n" +
		"/=======|  ",
	" +------+\n" +
		" |      |\n" +
		" |      |\n" +
		"(O)     |\n" +
		"/|\\     |\n" +
		"/       |\n" +
		"        |\n" +
		"        |\n" +
		"/=======|  ",
	" +------+\n" +
		" |      |\n" +
		" |      |\n" +
		"(O)     |\n" +
		"/|\\     |\n" +
		"/ \\     |\n" +
		"        |\n" +
		"        |\n" +
		"/=======|  ",
}

var gulliotine = [7]string{
	"|\n" +
		"|\n" +
		"|\n" +
		"|\n" +
		"|\n" +
		"|\n" +
		"|\n" +
		"|\n" +
		"|=======",
	"|       |\n" +
		"|       |\n" +
		"|       |\n" +
		"|       |\n" +
		"|       |\n" +
		"|       |\n" +
		"|       |\n" +
		"|       |\n" +
		"|=======|",
	"|=======|\n" +
		"|       |\n" +
		"|       |\n" +
		"|       |\n" +
		"|       |\n" +
		"|       |\n" +
		"|       |\n" +
		"|       |\n" +
		"|=======|",
	"|=======|\n" +
		"|       |\n" +
		"|       |\n" +
		"|       |\n" +
		"|       |\n" +
		"|       |\n" +
		"|       |\n" +
		"||-( )-||\n" +
		"|=======|",
	"|=======|\n" +
		"|       |\n" +
		"|       |\n" +
		"|       |\n" +
		"|       |\n" +
		"|       |\n" +
		"|/-----\\|\n" +
		"||-( )-||\n" +
		"|=======|",
	"|=======|\n" +
		"||     /|\n" +
		"||____/ |\n" +
		"|       |\n" +
		"|       |\n" +
		"|       |\n" +
		"|/-----\\|\n" +
		"||-( )-||\n" +
		"|=======|",

	"|=======|\n" +
		"||     /|\n" +
		"||____/ |\n" +
		"|       |\n" +
		"|       |\n" +
		"|       |\n" +
		"|/-----\\|\n" +
		"||-(O)-||\n" +
		"|=======|",
}

var secretWords = []string{
	"FATE", "HOPE", "HUMAN", "PAIN", "SADNESS", "ANGER",
	"COMFORT", "IDENTITY", "EMPATHY", "ANDROID", "DESPAIR",
	"DOUBTS", "PRISONER",
}
