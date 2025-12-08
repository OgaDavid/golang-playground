package main

import "fmt"

// Deal a pack of 12 cards evenly between 4 players, Player 1, Player 2, Player 3 and Player 4.

// Write a function DealAPackOfCards() that takes the argument, deck, as a slice of int and prints the desired output.

// Each player must be printed in a different line.

func DealAPackOfCards(deck []int) {
	var player1, player2, player3, player4 []int
	for _, n := range deck {
		if len(player1) == 3 && len(player2) == 3 {
			player3 = append(player3, n)
		} else if len(player2) == 3 && len(player3) == 0 {
			player4 = append(player4, n)
		} else if len(player1) == 3 {
			player2 = append(player2, n)
		}
		player1 = append(player1, n)
	}

	fmt.Printf("Player 1: %v", player1)
	fmt.Printf("Player 2: %v", player2)
	fmt.Printf("Player 3: %v", player3)
	fmt.Printf("Player 4: %v", player4)
}
