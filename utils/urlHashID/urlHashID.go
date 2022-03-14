package urlhashid

import (
	"math/rand"
	"time"
)

type IUrlHashID interface {
	generateHash() string
}

const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

func GenerateHash() string {
	rand.Seed(time.Now().UnixNano())
	min := 5
	max := 10
	length := rand.Intn(max-min) + min
	return String(length)
}
