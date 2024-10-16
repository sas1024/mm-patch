package patch

import (
	crand "crypto/rand"
	"math/rand"
	"time"
)

var (
	numbersAlphabet                = "0123456789"
	numbersAndLettersAlphabet      = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbersAndLowerLettersAlphabet = "0123456789abcdefghijklmnopqrstuvwxyz"
	rnd                            = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func GenerateRandomNumbers(length int) string {
	return generateRandomAlphabet(length, numbersAlphabet)
}

func GenerateRandomString(length int) string {
	return generateRandomAlphabet(length, numbersAndLettersAlphabet)
}

func GenerateRandomLowerString(length int) string {
	return generateRandomAlphabet(length, numbersAndLowerLettersAlphabet)
}

func RandomElement[T comparable](slice []T) T {
	return slice[rnd.Intn(len(slice))]
}

func generateRandomAlphabet(length int, alphabet string) string {
	var bytes = make([]byte, length)

	_, err := crand.Read(bytes)
	if err != nil {
		for k := range bytes {
			bytes[k] = alphabet[rnd.Intn(len(alphabet))]
		}
	} else {
		for k, v := range bytes {
			bytes[k] = alphabet[v%byte(len(alphabet))]
		}
	}

	return string(bytes)
}
