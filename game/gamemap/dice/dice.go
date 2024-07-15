package dice

import (
	"crypto/rand"
	"math/big"
)

func GetRandomInt(num int) int {
	x, _ := rand.Int(rand.Reader, big.NewInt(int64(num)))
	return int(x.Int64())
}

// GetDiceRoll returns an integer from 1 to the number
func GetDiceRoll(num int) int {
	x, _ := rand.Int(rand.Reader, big.NewInt(int64(num)))
	return int(x.Int64())
}

func GetRandomBetween(low, high int) int {
	var rnd int = -1
	for {
		rnd = GetDiceRoll(high)
		if rnd >= low {
			break
		}
	}
	return rnd
}
