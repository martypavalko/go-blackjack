package main

import "fmt"

//// Create a card struct
//// Function to return a string with the card value and suit
//// Create a deck struct
//// Function to build the deck
//// Function to build hand
// Function to output the player hand
// Function to output the hidden dealer hand
//// Function to add card to hands
// Function to deal card
// Function to replace suit value with unicode
// Function to evaluate the score of a hand
// Function to print Deck (FOR TESTING)
// Function to shuffle the deck
// Main function to play the game

type Card struct {
	Value string
	Suit  string
}

func printCard(c Card) (card string) {
	return c.Value + c.Suit
}

type Deck struct {
	currentCard int
	cards       [52]Card
}

func buildDeck() Deck {
	values := [13]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	suits := [4]string{"C", "S", "D", "H"}
	cards := [52]Card{}
	var deck Deck

	n := 0
	for i := 0; i < len(suits); i++ {
		for j := 0; j < len(values); j++ {
			cards[n].Value = values[j]
			//fmt.Printf("%v, %v, %v\n", i, j, j*i) // DEBUG
			cards[n].Suit = suits[i]
			//fmt.Printf("%v, %v, %v\n", i, j, n) // DEBUG
			n++
		}
	}

	deck.cards = cards
	return deck
}

func printDeck(deck Deck) {
	for i := range deck.cards {
		fmt.Printf("%v \t %v\n", deck.cards[i].Value, deck.cards[i].Suit)
	}
	fmt.Printf("Deck has %v cards\n", len(deck.cards))
}

type Hand struct {
	owner      string
	numOfCards int
	card       [12]Card
}

func createHand(owner string) Hand {
	var h Hand
	h.owner = owner
	return h
}

func addCard(h Hand, c Card) Hand {
	if h.numOfCards < 11 {
		h.card[h.numOfCards] = c
		h.numOfCards++
	}
	return h
}

func dealCard(h Hand, d Deck) Hand {

	h = addCard(h, d.cards[d.currentCard])

	return h
}

func main() {
	printDeck(buildDeck())
}
