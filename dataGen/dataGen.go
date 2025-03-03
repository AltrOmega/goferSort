package dataGen

import (
	"math/rand"
	"time"
)

func GenerateRandomSlice(size, min, max int) []int {
	rand.Seed(time.Now().UnixNano())
	randomSlice := make([]int, size)

	for i := 0; i < size; i++ {
		randomSlice[i] = rand.Intn(max-min+1) + min
	}

	return randomSlice
}
