package randgen

import (
	"math/rand"
	"time"
)

// StringGenerator is the interface that describes a
// string generator object.
type StringGenerator interface {
	GenerateFromSource(source string, length int) string
}

// stringGen is the object that satisfies the StringGenerator
// interface.
type stringGen struct{}

// NewStringGenerator returns a new string generator object that
// implements the StringGenerator interface.
func NewStringGenerator() StringGenerator {
	return &stringGen{}
}

// LowercaseStringAlphabetSource is a string generator source string that contains
// only lowercase alphabets.
var LowercaseStringAlphabetSource string = "qwertyuiopasdfghjklzxcvbnm"

// GenerateFromSource generates a random string with the specified
// length from source.
func (sg *stringGen) GenerateFromSource(source string, length int) string {
	rand.Seed(time.Now().Unix())
	newGeneratedRune := make([]rune, length)
	for i := 0; i < length; i++ {
		newGeneratedRune[i] = rune(source[rand.Intn(len(source))])
	}
	return string(newGeneratedRune)
}
