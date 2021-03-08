package randgen

// StringGenerator is the interface that describes a
// string generator object.
type StringGenerator interface {
}

// stringGen is the object that satisfies the StringGenerator
// interface.
type stringGen struct{}

// NewStringGenerator returns a new string generator object that
// implements the StringGenerator interface.
func NewStringGenerator() StringGenerator {
	return &stringGen{}
}
