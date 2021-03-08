package randgen

// NumberGenerator is the interface that describes a number generator
// object.
type NumberGenerator interface{}

// numberGen is the object that satisfies the NumberGenerator interface.
type numberGen struct{}

// NewNumberGenerator returns a new number generator object that implements
// the NumberGenerator interface.
func NewNumberGenerator() NumberGenerator {
	return &numberGen{}
}
