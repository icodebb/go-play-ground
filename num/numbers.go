/**
 * Play numbers.
 */

package num

import (
	"math/rand"

	log "github.com/sirupsen/logrus"
)

// Fibonacci produces Fibonacci numbers.
func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// RandomInt returns an int >= min, < max
func RandomInt(min, max int) int {
	// rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

// RandomString generates a random string of A-Z chars with len = l
func RandomString(len int) string {
	// rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(RandomInt(65, 90))
	}
	return string(bytes)
}

func NumTest() {
	log.Infof("Random int:%v, string:%s", RandomInt(1, 10), RandomString(8))
}
