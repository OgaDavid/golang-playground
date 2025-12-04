package main

import (
	"fmt"
)

// Deal a pack of 12 cards evenly between 4 players, Player 1, Player 2, Player 3 and Player 4.

// Write a function DealAPackOfCards() that takes the argument, deck, as a slice of int and prints the desired output.

// Each player must be printed in a different line.

func DealAPackOfCards(deck []int) {
	for i := 0; i < len(deck); i += 3 {
		playerNumber := i/3 + 1
		fmt.Printf("Player %d: %v, %v, %v\n", playerNumber, deck[i], deck[i+1], deck[i+2])
	}
}
