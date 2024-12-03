package main

import "slices"

type hand string

const (
	// Pocket pairs
	handAA hand = "AA"
	handKK hand = "KK"
	handQQ hand = "QQ"
	handJJ hand = "JJ"
	handTT hand = "TT"
	hand99 hand = "99"
	hand88 hand = "88"
	hand77 hand = "77"
	hand66 hand = "66"
	hand55 hand = "55"
	hand44 hand = "44"
	hand33 hand = "33"
	hand22 hand = "22"

	// Suited hands
	handAKs hand = "AKs"
	handAQs hand = "AQs"
	handAJs hand = "AJs"
	handATs hand = "ATs"
	handA9s hand = "A9s"
	handA8s hand = "A8s"
	handA7s hand = "A7s"
	handA6s hand = "A6s"
	handA5s hand = "A5s"
	handA4s hand = "A4s"
	handA3s hand = "A3s"
	handA2s hand = "A2s"
	handKQs hand = "KQs"
	handKJs hand = "KJs"
	handKTs hand = "KTs"
	handK9s hand = "K9s"
	handK8s hand = "K8s"
	handK7s hand = "K7s"
	handK6s hand = "K6s"
	handK5s hand = "K5s"
	handK4s hand = "K4s"
	handK3s hand = "K3s"
	handK2s hand = "K2s"
	handQJs hand = "QJs"
	handQTs hand = "QTs"
	handQ9s hand = "Q9s"
	handQ8s hand = "Q8s"
	handQ7s hand = "Q7s"
	handQ6s hand = "Q6s"
	handQ5s hand = "Q5s"
	handQ4s hand = "Q4s"
	handQ3s hand = "Q3s"
	handQ2s hand = "Q2s"
	handJTs hand = "JTs"
	handJ9s hand = "J9s"
	handJ8s hand = "J8s"
	handJ7s hand = "J7s"
	handJ6s hand = "J6s"
	handJ5s hand = "J5s"
	handJ4s hand = "J4s"
	handJ3s hand = "J3s"
	handJ2s hand = "J2s"
	handT9s hand = "T9s"
	handT8s hand = "T8s"
	handT7s hand = "T7s"
	handT6s hand = "T6s"
	handT5s hand = "T5s"
	handT4s hand = "T4s"
	handT3s hand = "T3s"
	handT2s hand = "T2s"
	hand98s hand = "98s"
	hand97s hand = "97s"
	hand96s hand = "96s"
	hand95s hand = "95s"
	hand94s hand = "94s"
	hand93s hand = "93s"
	hand92s hand = "92s"
	hand87s hand = "87s"
	hand86s hand = "86s"
	hand85s hand = "85s"
	hand84s hand = "84s"
	hand83s hand = "83s"
	hand82s hand = "82s"
	hand76s hand = "76s"
	hand75s hand = "75s"
	hand74s hand = "74s"
	hand73s hand = "73s"
	hand72s hand = "72s"
	hand65s hand = "65s"
	hand64s hand = "64s"
	hand63s hand = "63s"
	hand62s hand = "62s"
	hand54s hand = "54s"
	hand53s hand = "53s"
	hand52s hand = "52s"
	hand43s hand = "43s"
	hand42s hand = "42s"
	hand32s hand = "32s"

	// Offsuit hands
	handAKo hand = "AKo"
	handAQo hand = "AQo"
	handAJo hand = "AJo"
	handATo hand = "ATo"
	handA9o hand = "A9o"
	handA8o hand = "A8o"
	handA7o hand = "A7o"
	handA6o hand = "A6o"
	handA5o hand = "A5o"
	handA4o hand = "A4o"
	handA3o hand = "A3o"
	handA2o hand = "A2o"
	handKQo hand = "KQo"
	handKJo hand = "KJo"
	handKTo hand = "KTo"
	handK9o hand = "K9o"
	handK8o hand = "K8o"
	handK7o hand = "K7o"
	handK6o hand = "K6o"
	handK5o hand = "K5o"
	handK4o hand = "K4o"
	handK3o hand = "K3o"
	handK2o hand = "K2o"
	handQJo hand = "QJo"
	handQTo hand = "QTo"
	handQ9o hand = "Q9o"
	handQ8o hand = "Q8o"
	handQ7o hand = "Q7o"
	handQ6o hand = "Q6o"
	handQ5o hand = "Q5o"
	handQ4o hand = "Q4o"
	handQ3o hand = "Q3o"
	handQ2o hand = "Q2o"
	handJTo hand = "JTo"
	handJ9o hand = "J9o"
	handJ8o hand = "J8o"
	handJ7o hand = "J7o"
	handJ6o hand = "J6o"
	handJ5o hand = "J5o"
	handJ4o hand = "J4o"
	handJ3o hand = "J3o"
	handJ2o hand = "J2o"
	handT9o hand = "T9o"
	handT8o hand = "T8o"
	handT7o hand = "T7o"
	handT6o hand = "T6o"
	handT5o hand = "T5o"
	handT4o hand = "T4o"
	handT3o hand = "T3o"
	handT2o hand = "T2o"
	hand98o hand = "98o"
	hand97o hand = "97o"
	hand96o hand = "96o"
	hand95o hand = "95o"
	hand94o hand = "94o"
	hand93o hand = "93o"
	hand92o hand = "92o"
	hand87o hand = "87o"
	hand86o hand = "86o"
	hand85o hand = "85o"
	hand84o hand = "84o"
	hand83o hand = "83o"
	hand82o hand = "82o"
	hand76o hand = "76o"
	hand75o hand = "75o"
	hand74o hand = "74o"
	hand73o hand = "73o"
	hand72o hand = "72o"
	hand65o hand = "65o"
	hand64o hand = "64o"
	hand63o hand = "63o"
	hand62o hand = "62o"
	hand54o hand = "54o"
	hand53o hand = "53o"
	hand52o hand = "52o"
	hand43o hand = "43o"
	hand42o hand = "42o"
	hand32o hand = "32o"
)

var allHands = [13 * 13]hand{
	handAA, handAKs, handAQs, handAJs, handATs, handA9s, handA8s, handA7s, handA6s, handA5s, handA4s, handA3s, handA2s,
	handAKo, handKK, handKQs, handKJs, handKTs, handK9s, handK8s, handK7s, handK6s, handK5s, handK4s, handK3s, handK2s,
	handAQo, handKQo, handQQ, handQJs, handQTs, handQ9s, handQ8s, handQ7s, handQ6s, handQ5s, handQ4s, handQ3s, handQ2s,
	handAJo, handKJo, handQJo, handJJ, handJTs, handJ9s, handJ8s, handJ7s, handJ6s, handJ5s, handJ4s, handJ3s, handJ2s,
	handATo, handKTo, handQTo, handJTo, handTT, handT9s, handT8s, handT7s, handT6s, handT5s, handT4s, handT3s, handT2s,
	handA9o, handK9o, handQ9o, handJ9o, handT9o, hand99, hand98s, hand97s, hand96s, hand95s, hand94s, hand93s, hand92s,
	handA8o, handK8o, handQ8o, handJ8o, handT8o, hand98o, hand88, hand87s, hand86s, hand85s, hand84s, hand83s, hand82s,
	handA7o, handK7o, handQ7o, handJ7o, handT7o, hand97o, hand87o, hand77, hand76s, hand75s, hand74s, hand73s, hand72s,
	handA6o, handK6o, handQ6o, handJ6o, handT6o, hand96o, hand86o, hand76o, hand66, hand65s, hand64s, hand63s, hand62s,
	handA5o, handK5o, handQ5o, handJ5o, handT5o, hand95o, hand85o, hand75o, hand65o, hand55, hand54s, hand53s, hand52s,
	handA4o, handK4o, handQ4o, handJ4o, handT4o, hand94o, hand84o, hand74o, hand64o, hand54o, hand44, hand43s, hand42s,
	handA3o, handK3o, handQ3o, handJ3o, handT3o, hand93o, hand83o, hand73o, hand63o, hand53o, hand43o, hand33, hand32s,
	handA2o, handK2o, handQ2o, handJ2o, handT2o, hand92o, hand82o, hand72o, hand62o, hand52o, hand42o, hand32o, hand22,
}

type handList []hand

func (l handList) Contains(h hand) bool {
	return slices.Contains(l, h)
}

type rank int

const (
	rankUTGStrong rank = iota
	rankUTGMiddle
	rankUTGWeak
	rankEP
	rankLJHJ
	rankCO
	rankBTN
	rankBB
	rankFold
)

type HandRankMap map[rank]handList

var handRankMapCache HandRankMap

func handRanks() HandRankMap {
	if len(handRankMapCache) != 0 {
		return handRankMapCache
	}

	m := HandRankMap{
		rankUTGStrong: {
			handAA, handKK, handQQ,
			handAKs,
			handAKo,
		},
		rankUTGMiddle: {
			handJJ, handTT, hand99,
			handAQs, handAJs, handATs, handKQs,
			handAQo,
		},
		rankUTGWeak: {
			hand88, hand77,
			handKJs, handQJs, handJTs,
			handAJo, handKQo,
		},
		rankEP: {
			hand66, hand55,
			handA9s, handA8s, handA7s, handA6s, handA5s, handA4s, handA3s, handA2s, handKTs, handK9s, handQTs, handT9s,
			handATo, handKJo,
		},
		rankLJHJ: {
			hand44, hand33, hand22,
			handQ9s, handJ9s, handT8s, hand98s,
			handA9o, handKTo, handQJo, handJTo,
		},
		rankCO: {
			handK8s, handK7s, handK6s, handK5s, handK4s, handK3s, handK2s, handQ8s, handQ7s, handQ6s, handJ8s, handJ7s, hand97s, hand87s, hand76s, hand65s,
			handA8o, handA7o, handK9o, handQTo, handQ9o, handJ9o, handT9o,
		},
		rankBTN: {
			handQ5s, handQ4s, handQ3s, handQ2s, handJ6s, handT7s, hand96s, hand86s, hand75s, hand64s, hand54s,
			handA6o, hand98o,
		},
		rankBB: {
			handJ5s, handJ4s, handJ3s, handJ2s, handT6s, handT5s, handT4s, handT3s, hand95s, hand85s, hand74s, hand63s, hand53s, hand43s,
			handA5o, handA4o, handA3o, handA2o, handK8o, handK7o, handK6o, handK5o, handQ8o, handQ7o, handJ8o, handT8o, hand97o, hand87o,
		},
	}

	var foldHands []hand
	for _, h := range allHands {
		var found bool
		for _, v := range m {
			if slices.Contains(v, h) {
				found = true
				break
			}
		}
		if !found {
			foldHands = append(foldHands, h)
		}
	}
	m[rankFold] = foldHands

	handRankMapCache = m
	return handRankMapCache
}
