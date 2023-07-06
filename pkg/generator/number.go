package generator

import (
	"math/rand"
	"time"
)

func GenerateInt(max int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Intn(max)
}

func GenerateFloat() float64 {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Float64()
}
