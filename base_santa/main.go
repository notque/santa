package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ppl := []string{"Vader", "Luke", "Leia", "Han", "Palpatine", "Chewie", "C3PO"}

	rand.Seed(time.Now().UnixNano())
	result, err := secretSanta(ppl)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		for _, pair := range result {
			fmt.Printf("(%s, %s)\n", pair[0], pair[1])
		}
	}
}

func secretSanta(ppl []string) ([][]string, error) {
	n := len(ppl)
	if n < 2 {
		return nil, fmt.Errorf("not enough people for Secret Santa")
	}

	shuffled := make([]string, n)
	copy(shuffled, ppl)
	rand.Shuffle(n, func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	// Check if the last person got themselves, if so, swap with someone else
	if shuffled[n-1] == ppl[n-1] {
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

