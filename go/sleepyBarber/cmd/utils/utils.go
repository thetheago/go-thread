package utils

import (
	"math/rand"
	"time"
)

func SleepRandom() {
	r := rand.Intn(10)
	time.Sleep(time.Duration(r) * time.Second)
}
