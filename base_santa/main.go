package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// List of participants
	participants := []string{"Vader", "Luke", "Leia", "Han", "Palpatine", "Chewie", "C3PO"}

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate Secret Santa pairs
	result, err := secretSanta(participants)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		for _, pair := range result {
			fmt.Printf("(%s, %s)\n", pair[0], pair[1])
		}
	}
}

// secretSanta generates Secret Santa pairs from a list of participants.
func secretSanta(participants []string) ([][]string, error) {
	n := len(participants)
	if n < 2 {
		return nil, fmt.Errorf("not enough people for Secret Santa")
	}

	// Shuffle the participants
	shuffled := make([]string, n)
	copy(shuffled, participants)
	rand.Shuffle(n, func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	// Check if the last person got themselves, if so, swap with someone else
	if shuffled[n-1] == participants[n-1] {
		shuffled[n-1], shuffled[0] = shuffled[0], shuffled[n-1]
	}

	// Pair each person with the next person in the shuffled list
	pairs := make([][]string, n)
	for i := 0; i < n; i++ {
		receiver := shuffled[(i+1)%n]
		pairs[i] = []string{shuffled[i], receiver}
	}

	return pairs, nil
}

