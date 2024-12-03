package main

import (
	"fmt"
	"math/rand"
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

		hrs := handRanks()
		if hrs[input].Contains(allHands[questionIdx]) {
			fmt.Println("Correct!")
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
}
