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

// StringLowercaseAlphabetsSource is a string generator source string that contains
// only lowercase alphabets.
var StringLowercaseAlphabetsSource string = "qwertyuiopasdfghjklzxcvbnm"

// StringUppercaseAlphabetsSource is a string generator source string that contains
// only uppercase alphabets.
var StringUppercaseAlphabetsSource string = "QWERTYUIOPASDFGHJKLZXCVBNM"

// StringAlphabetsSource is a string generator source that contains both uppercase and
// lowercase alphabets only.
var StringAlphabetsSource string = StringLowercaseAlphabetsSource + StringUppercaseAlphabetsSource

// StringSymbolsSource is a string generator source that contains
// symbols only.
var StringSymbolsSource string = "~`!@#$%^&*()_+-={}[]|\\;:'\",.<>/?"

// StringNumbericSoure is a string generator source that contains
// numberic characters only.
var StringNumbericSoure string = "0123456789"

// StringLowercaseAlphaNumericSource is a string generator source that contains
// lowercase alph-numeric characters only.
var StringLowercaseAlphaNumericSource string = StringLowercaseAlphabetsSource + StringNumbericSoure

// StringUppercaseAlphaNumericSource is a string generator source that contains
// uppercase alph-numeric characters only.
var StringUppercaseAlphaNumericSource string = StringUppercaseAlphabetsSource + StringNumbericSoure

// StringAlphaNumericSource is a string generator source that contains
// all cases alph-numeric characters.
var StringAlphaNumericSource string = StringLowercaseAlphaNumericSource + StringUppercaseAlphaNumericSource

// NewStringGenerator returns a new string generator object that
// implements the StringGenerator interface.
func NewStringGenerator() StringGenerator {
	return &stringGen{}
}

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
