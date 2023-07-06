package utils

import (
	"math/rand"
	"time"
)

func RandomInt(max int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Intn(max)
}

func RandomFloat64() float64 {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Float64()
}
