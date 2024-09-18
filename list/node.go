package list

import (
	"math/rand"
        "cmp"
)

type Node[T cmp.Ordered] struct {
	value T
	next  []*Node[T] // forward next pointers (# is randomly generated)
}

// Node constructor
func NewNode[T cmp.Ordered](value T) (*Node[T], int) {
	height := RandomLevel() + 1
	next := make([]*Node[T], height)
	node := &Node[T]{value: value, next: next}
	return node, height
}

// Generate random level
// Each successive level generation probability is (1/2)^(# coin flips)
func RandomLevel() int {
	var level int

	// Every time the proverbial "coin flip" returns true, track increments
	// as long as they are valid (1/2 * 1/2 * ....)
	for rand.Intn(2) == 0 && level < SkipLevels-1 {
		level++
	}

	return level
}
