package utils

import (
	"math/rand"
	"strconv"
)

// Generates a random numeric string of given length
func GenerateRandomNumeric(len int) string {
	var randSeq = make([]byte, len)
	for i, _ := range randSeq {
		randSeq[i] = []byte(strconv.Itoa(rand.Intn(10)))[0]
	}

	return string(randSeq)
}
