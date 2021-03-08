package randgen

import (
	"math/rand"
	"strconv"
	"time"
)

// NumberGenerator is the interface that describes a number generator
// object.
type NumberGenerator interface {
	GenerateFromSource(source, length int) int
}

// numberGen is the object that satisfies the NumberGenerator interface.
type numberGen struct{}

// NumberSource is the source for number generator that contains
// all the numbers.
var NumberSource int = 1234567890

// NewNumberGenerator returns a new number generator object that implements
// the NumberGenerator interface.
func NewNumberGenerator() NumberGenerator {
	return &numberGen{}
}

// GenerateFromSource generates random numbers with a specified length with
// the numbers gotten from the source.
func (ng *numberGen) GenerateFromSource(source, length int) int {
	rand.Seed(time.Now().Unix())
	sourceStr := strconv.Itoa(source)
	result := make([]rune, length)
	for i := 0; i < length; i++ {
		randomVal := rune(sourceStr[rand.Intn(len(sourceStr)-1)])
		if i == 0 && string(randomVal) == "0" {
			i = 0
			continue
		}
		result[i] = randomVal
	}
	resultInt, _ := strconv.Atoi(string(result))
	return resultInt
}
