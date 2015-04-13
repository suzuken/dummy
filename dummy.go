package dummy

import (
	"math/rand"
)

var (
	letters = []rune("abcdefghijklmnopqrstuvwxyz")
	numbers = []rune("0123456789")
)

func Gen(r []rune, n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = r[rand.Intn(n) % len(r)]
	}
	return string(b)
}

func String(n int) string {
	return Gen(letters, n)
}

func Int(n int) string {
	return Gen(numbers, n)
}
