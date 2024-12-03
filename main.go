package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

var answerOptions = []string{
	rankUTGStrong: "Strong hand for UTG",
	rankUTGMiddle: "Middle hand for UTG",
	rankUTGWeak:   "Weak hand for UTG",
	rankEP:        "Hand for EP",
	rankLJHJ:      "Hand for LJ/HJ",
	rankCO:        "Hand for CO",
	rankBTN:       "Hand for BTN",
	rankBB:        "Hand for BB in the case the BTN raises",
	rankFold:      "Hand to fold",
}

func main() {
	fmt.Println("Welcome to the hand rank quiz!")
	fmt.Println("If you want to finish the quiz, enter 'q'")
	fmt.Println()

	countQuestions := 0
	countCorrect := 0
	countQuestionsCh := make(chan int)
	countCorrectCh := make(chan int)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			var input string

			questionIdx := rand.Intn(13 * 13)

			fmt.Printf("The hand is %s\n", allHands[questionIdx])
			fmt.Println("What is the rank of this hand?")

			for i, v := range answerOptions {
				fmt.Printf("%d. %s\n", i, v)
			}

			var answer rank
			for {
				fmt.Print("> ")
				fmt.Scan(&input)

				if input == "q" {
					quit <- syscall.SIGINT
					return
				}

				a, err := strconv.Atoi(input)
				if err != nil || a < 0 || a >= len(answerOptions) {
					fmt.Println("Invalid input. Please enter a number between 0 and 8 or 'q' to quit.")
					continue
				}
				answer = rank(a)
				break
			}

			countQuestionsCh <- 1

			hrs := handRanks()
			if hrs[answer].Contains(allHands[questionIdx]) {
				fmt.Println("Correct!")
				countCorrectCh <- 1
			} else {
				fmt.Println("Incorrect.")

				var correctAnswerMsg string
				for i, v := range hrs {
					if v.Contains(allHands[questionIdx]) {
						correctAnswerMsg = fmt.Sprintf("The correct answer is: %s", answerOptions[i])
						break
					}
				}
				fmt.Println(correctAnswerMsg)
			}

			fmt.Println()
			fmt.Println("--------------------")
			fmt.Println()
		}
	}()

	for {
		select {
		case <-quit:
			fmt.Println()
			fmt.Println("Quiz finished!")
			fmt.Printf("Result: %d/%d\n", countCorrect, countQuestions)
			return
		case c := <-countQuestionsCh:
			countQuestions += c
		case c := <-countCorrectCh:
			countCorrect += c
		}
	}
}
