package pick

import (
	"errors"
	"fmt"
	"math/rand"
)

// ErrEmptyCollection is when a slice has no elements.
var ErrEmptyCollection = errors.New("empty collection")

// Picker provides a shared set of functionality for picking elements from a slice.
type Picker struct {
	Rand func(int) int
}

// NewPicker creates a new instance of a picker.
func NewPicker() *Picker {
	return &Picker{
		Rand: rand.Intn,
	}
}

// GetOne returns a single random element from a slice. The function to select items may be customized.
func GetOne[E any](p *Picker, src []E) (E, error) {
	if len(src) == 0 {
		return *new(E), fmt.Errorf("%w: %+v", ErrEmptyCollection, src)
	}

	idx := p.Rand(len(src))

	return src[idx], nil
}
