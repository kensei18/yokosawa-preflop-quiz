package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
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
	fmt.Println()

	countQuestions := 0
	countCorrect := 0
	countQuestionsCh := make(chan int)
	countCorrectCh := make(chan int)

	go func() {
		for {
			var input rank

			questionIdx := rand.Intn(13 * 13)

			fmt.Printf("The hand is %s\n", allHands[questionIdx])
			fmt.Println("What is the rank of this hand?")

			for i, v := range answerOptions {
				fmt.Printf("%d. %s\n", i, v)
			}

			fmt.Print("> ")
			fmt.Scan(&input)

			countQuestionsCh <- 1

			hrs := handRanks()
			if hrs[input].Contains(allHands[questionIdx]) {
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

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-quit:
			fmt.Println()
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
