package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// List of participants
	participants := []string{"Vader", "Luke", "Leia", "Han", "Palpatine", "Chewie", "C3PO"}

	// Forbidden pairs
	noGifts := map[string]map[string]bool{
		"Palpatine": {"Vader": true, "Luke": true},
		"Chewie":    {"Palpatine": true},
	}

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate Secret Santa pairs with forbidden pairs consideration
	result, err := secretSantaWithRestrictions(participants, noGifts)
	if err != nil {
		fmt.Println("No solution found.")
	} else {
		for _, pair := range result {
			fmt.Printf("%s --> %s\n", pair[0], pair[1])
		}
	}
}

// secretSantaWithRestrictions generates Secret Santa pairs with forbidden pairs consideration.
func secretSantaWithRestrictions(participants []string, noGifts map[string]map[string]bool) ([][]string, error) {
	for attempts := 0; attempts < 1000; attempts++ {
		// Shuffle the participants
		shuffled := shuffle(participants)
		valid := true

		// Check for valid pairs
		for i, giver := range shuffled {
			receiver := shuffled[(i+1)%len(shuffled)]
			if giver == receiver || noGifts[giver][receiver] {
				valid = false
				break
			}
		}

		// If all pairs are valid, return the result
		if valid {
			pairs := make([][]string, len(shuffled))
			for i, giver := range shuffled {
				receiver := shuffled[(i+1)%len(shuffled)]
				pairs[i] = []string{giver, receiver}
			}
			return pairs, nil
		}
	}

	// If no valid solution is found after a number of attempts, return an error
	return nil, fmt.Errorf("no valid solution found")
}

// shuffle randomly shuffles a slice of strings.
func shuffle(slice []string) []string {
	shuffled := make([]string, len(slice))
	copy(shuffled, slice)
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	return shuffled
}
