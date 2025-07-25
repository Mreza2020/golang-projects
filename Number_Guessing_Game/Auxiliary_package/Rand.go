package Auxiliary_package

import (
	"math/rand"
	"time"
)

func Rand(number int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(number-1+1) + 1
}
