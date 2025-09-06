package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateMultiplier(rtp float64) float64 {
	u := rand.Float64()
	t1 := 1 - rtp
	t2 := 1 - (rtp / 10000.0)
	if u < t1 {
		return 1.0
	} else if u > t2 {
		return 10000.0
	}
	return rtp / (1 - u)
}
