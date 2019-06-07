package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)


type player struct {
	name  string
	cards []card
}

type card struct {
	suit		string
	value		int
	ofAKind 	int
	straight	int
}

var (
	deck 		[]card
	usedDeck	[]card
	players		[]player
)



func main() {
	rand.Seed(time.Now().UnixNano())
	numberOfPlayers := getNumberOfPlayers()
	buildDeck()
	fmt.Println("Number of Players: ", numberOfPlayers)
	fmt.Println("\nLets begin playing rummy.....")

	deal(numberOfPlayers)

	for _, player := range players {
		checkHand(&player.cards)
		printHand(player)
	}

	fmt.Println("\nPlayers' hands: ", players)
}


func getNumberOfPlayers() int {
	if len(os.Args) != 2 {
		fmt.Println("Specify a number of players. Usage:\n", os.Args[0], "<No of Players>")
		os.Exit(3)
	}

	numberOfPlayers, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Please give a valid number of players")
	}

	if numberOfPlayers > 5 || numberOfPlayers < 1 {
		fmt.Println("Usage: Max Number of Players is 5, Minimum is 1")
		os.Exit(3)
	}

	return numberOfPlayers
}


func buildDeck() {
	suits := []string{"Spades", "Clubs", "Diamonds", "Hearts"}
	for value := 2; value < 15; value++ {
		for _, suit := range suits {
			deck = append(deck, card{suit: suit, value: value})
		}
	}


	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })
}


func deal(numberOfPlayers int) {
	for playerNumber := 1; playerNumber <= numberOfPlayers; playerNumber++ {
		name := fmt.Sprintf("Player %d", playerNumber)
		cardsInHand := 7
		if playerNumber == 1 {
			cardsInHand = 8
		}
		startingHand := deck[:cardsInHand]
		deck = deck[cardsInHand:]
		

		players = append(players, player{name: name, cards: startingHand})
	}
}

func checkHand(cards *[]card) {
	hand := *cards
	for i, card := range hand {
		for j, card2 := range hand {
			if i == j {
				continue
			}
			if card2.suit == card.suit {
				diff := card2.value - card.value
				if diff < 5 && diff > -5 {
					(*cards)[i].straight++
				}
			} else {
				if card2.value == card.value {
					(*cards)[i].ofAKind++
				}
			}
		}
	}
}


func printHand(player player) {
	var hand string

	for _, card := range player.cards {
		hand += strconv.Itoa(card.value) + " of " + card.suit + ", "
	}

	hand = hand[:len(hand)-2] + "."

	fmt.Println(player.name, " holds:")
	fmt.Println(hand)
}
