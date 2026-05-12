package utils

import (
	"math/rand"
	"time"
)

func SleepBarberCuttingHair() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	segundos := r.Intn(6) + 6

	time.Sleep(time.Duration(segundos) * time.Second)
}

func SleepBarberCustomerGap() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	segundos := r.Intn(1) + 5

	time.Sleep(time.Duration(segundos) * time.Second)
}

func SleepRandomLongLong() {
	r := rand.Intn(20)
	time.Sleep(time.Duration(r) * time.Second)
}

func SleepRandomLong() {
	r := rand.Intn(10)
	time.Sleep(time.Duration(r) * time.Second)
}

func SleepRandomMedium() {
	r := rand.Intn(6)
	time.Sleep(time.Duration(r) * time.Second)
}

func SleepRandomShort() {
	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)
}
