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
	suit	string
	value	int
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
