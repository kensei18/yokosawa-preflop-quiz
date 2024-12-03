package main

import (
	"fmt"
	"math/rand"
	"slices"
)

const (
	strengthUTGStrong = iota
	strengthUTGMiddle
	strengthUTGWeak
	strengthEP
	strengthLJHJ
	strengthCO
	strengthBTN
	strengthBB
	strengthFold
)

var handList = [][]string{
	{"AA", "AKs", "AQs", "AJs", "ATs", "A9s", "A8s", "A7s", "A6s", "A5s", "A4s", "A3s", "A2s"},
	{"AKo", "KK", "KQs", "KJs", "KTs", "K9s", "K8s", "K7s", "K6s", "K5s", "K4s", "K3s", "K2s"},
	{"AQo", "KQo", "QQ", "QJs", "QTs", "Q9s", "Q8s", "Q7s", "Q6s", "Q5s", "Q4s", "Q3s", "Q2s"},
	{"AJo", "KJo", "QJo", "JJ", "JTs", "J9s", "J8s", "J7s", "J6s", "J5s", "J4s", "J3s", "J2s"},
	{"ATo", "KTo", "QTo", "JTo", "TT", "T9s", "T8s", "T7s", "T6s", "T5s", "T4s", "T3s", "T2s"},
	{"A9o", "K9o", "Q9o", "J9o", "T9o", "99", "98s", "97s", "96s", "95s", "94s", "93s", "92s"},
	{"A8o", "K8o", "Q8o", "J8o", "T8o", "98o", "88", "87s", "86s", "85s", "84s", "83s", "82s"},
	{"A7o", "K7o", "Q7o", "J7o", "T7o", "97o", "87o", "77", "76s", "75s", "74s", "73s", "72s"},
	{"A6o", "K6o", "Q6o", "J6o", "T6o", "96o", "86o", "76o", "66", "65s", "64s", "63s", "62s"},
	{"A5o", "K5o", "Q5o", "J5o", "T5o", "95o", "85o", "75o", "65o", "55", "54s", "53s", "52s"},
	{"A4o", "K4o", "Q4o", "J4o", "T4o", "94o", "84o", "74o", "64o", "54o", "44", "43s", "42s"},
	{"A3o", "K3o", "Q3o", "J3o", "T3o", "93o", "83o", "73o", "63o", "53o", "43o", "33", "32s"},
	{"A2o", "K2o", "Q2o", "J2o", "T2o", "92o", "82o", "72o", "62o", "52o", "42o", "32o", "22"},
}

var handRangeMap = map[int][][]int{
	strengthUTGStrong: {
		{0, 0}, {0, 1}, {1, 0}, {1, 1}, {2, 2},
	},
	strengthUTGMiddle: {
		{0, 2}, {0, 3}, {0, 4}, {1, 2}, {0, 2}, {3, 3}, {4, 4}, {5, 5},
	},
	strengthUTGWeak: {
		{1, 3}, {2, 1}, {2, 3}, {3, 0}, {3, 4}, {6, 6}, {7, 7},
	},
	strengthEP: {
		{0, 5}, {0, 6}, {0, 7}, {0, 8}, {0, 9}, {0, 10}, {0, 11}, {0, 12}, {1, 4}, {1, 5}, {2, 4}, {3, 1}, {4, 0}, {4, 5}, {8, 8}, {9, 9},
	},
	strengthLJHJ: {
		{2, 5}, {3, 2}, {3, 5}, {4, 1}, {4, 3}, {4, 6}, {5, 0}, {5, 6}, {10, 10}, {11, 11}, {12, 12},
	},
	strengthCO: {
		{1, 6}, {1, 7}, {1, 8}, {1, 9}, {1, 10}, {1, 11}, {1, 12}, {2, 6}, {2, 7}, {2, 8}, {3, 6}, {3, 7}, {4, 2}, {5, 1}, {5, 2}, {5, 3}, {5, 4}, {5, 7}, {6, 0}, {6, 7}, {7, 0}, {7, 8}, {8, 9},
	},
	strengthBTN: {
		{2, 9}, {2, 10}, {2, 11}, {2, 12}, {3, 8}, {4, 7}, {5, 8}, {6, 5}, {6, 8}, {7, 9}, {8, 0}, {8, 10}, {9, 10},
	},
	strengthBB: {
		{3, 9}, {3, 10}, {3, 11}, {3, 12}, {4, 8}, {4, 9}, {4, 10}, {4, 11}, {5, 9}, {6, 1}, {6, 2}, {6, 3}, {6, 4}, {6, 9}, {7, 1}, {7, 2}, {7, 5}, {7, 6}, {7, 10}, {8, 1}, {8, 11}, {9, 0}, {9, 1}, {9, 11}, {10, 0}, {10, 11}, {11, 0}, {12, 0},
	},
}

var handRangeAnswerOptions = map[int]string{
	strengthUTGStrong: "Strong hand for UTG",
	strengthUTGMiddle: "Middle hand for UTG",
	strengthUTGWeak:   "Weak hand for UTG",
	strengthEP:        "Hand for EP",
	strengthLJHJ:      "Hand for LJ/HJ",
	strengthCO:        "Hand for CO",
	strengthBTN:       "Hand for BTN",
	strengthBB:        "Hand for BB in the case the BTN raises",
	strengthFold:      "Hand to fold",
}

func main() {
	answerOptions := make([]string, 0, len(handRangeAnswerOptions))
	for i, v := range handRangeAnswerOptions {
		answerOptions = append(answerOptions, fmt.Sprintf("%d: %s", i, v))
	}
	slices.Sort(answerOptions)

	for {
		var input int

		a := rand.Intn(13)
		b := rand.Intn(13)

		fmt.Printf("The hand is %s\n", handList[a][b])
		fmt.Println("What is the strength of this hand?")

		for _, v := range answerOptions {
			fmt.Println(v)
		}

		fmt.Print("> ")
		fmt.Scan(&input)

		if slices.ContainsFunc(handRangeMap[input], func(v []int) bool {
			return v[0] == a && v[1] == b
		}) {
			fmt.Println("Correct!")
		} else {
			fmt.Println("Incorrect.")

			var correctAnswerMsg string
			for i, v := range handRangeMap {
				if i == strengthFold {
					continue
				}
				if slices.ContainsFunc(v, func(v []int) bool {
					return v[0] == a && v[1] == b
				}) {
					correctAnswerMsg = fmt.Sprintf("The correct answer is: %s", handRangeAnswerOptions[i])
					break
				}
			}
			if correctAnswerMsg == "" {
				correctAnswerMsg = "The correct answer is: Hand to fold"
			}
			fmt.Println(correctAnswerMsg)
		}

		fmt.Println()
		fmt.Println("--------------------")
		fmt.Println()
	}
}
