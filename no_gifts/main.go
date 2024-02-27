package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ppl := []string{"Vader", "Luke", "Leia", "Han", "Palpatine", "Chewie", "C3PO"}
	noGifts := map[string]map[string]bool{
		"Palpatine": {"Vader": true, "Luke": true},
		"Chewie":    {"Palpatine": true},
	}

	rand.Seed(time.Now().UnixNano())
	result, err := secretSanta(ppl, noGifts)
	if err != nil {
		fmt.Println("No solution found.")
	} else {
		for _, pair := range result {
			fmt.Printf("%s --> %s\n", pair[0], pair[1])
		}
	}
}

func secretSanta(ppl []string, noGifts map[string]map[string]bool) ([][]string, error) {
	for attempts := 0; attempts < 1000; attempts++ {
		shuffled := shuffle(ppl)
		valid := true
		for i, giver := range shuffled {
			receiver := shuffled[(i+1)%len(shuffled)]
			if giver == receiver || noGifts[giver][receiver] {
				valid = false
				break
			}
		}
		if valid {
			pairs := make([][]string, len(shuffled))
			for i, giver := range shuffled {
				receiver := shuffled[(i+1)%len(shuffled)]
				pairs[i] = []string{giver, receiver}
			}
			return pairs, nil
		}
	}
	return nil, fmt.Errorf("no valid solution found")
}

func shuffle(slice []string) []string {
	shuffled := make([]string, len(slice))
	copy(shuffled, slice)
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	return shuffled
}
