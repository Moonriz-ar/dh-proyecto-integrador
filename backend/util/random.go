package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	// set the seed value for the random generator
	// normally the seed value is set to current time, converted to unit nano since rand.Seed() expect int64
	// this will make sure that everytime we run the code, the generated values will be different
	// if we don't call rand.Seed(), the random generator will behave like it is seeded by 1, and the generated values will be the same for every run
	rand.Seed(time.Now().UnixNano())
}

// RandomInt takes 2 int64 numbers min and max, returns a random int64 number between min and max
func RandomInt(min, max int64) int64 {
	// rand.Int63n(n) function returns a random integer between 0 and n-1
	return min + rand.Int63n(max-min+1)
}

// RandomString takes 1 int number, returns a random string with specified length
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}
