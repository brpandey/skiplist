package list

import (
	"iter"
)

// Go 1.23 now supports standardized iterators
// Values is an iterator over the value elements of sl
func (sl *SkipList) Values() iter.Seq[int] {
	return func(yield func(int) bool) {
		if sl == nil {
			return
		}
		cur := sl.head
		for cur.next[0] != nil { // loop through bottom row only
			cur = cur.next[0]
			if !yield(cur.value) {
				return
			}
		}
	}
}

// All returns an iterator over the elements of sl, including level data as a
// two-element tuple
func (sl *SkipList) All() iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		if sl == nil {
			return
		}
		cur := sl.head
		top := sl.height - 1

		for i := top; i >= 0; i-- {
			for cur.next[i] != nil { // loop through bottom row only
				cur = cur.next[i]
				if !yield(i, cur.value) {
					return
				}
			}

			cur = sl.head
		}
	}
}

// AllUnique returns an iterator over the unique elements of sl with the highest
// level they first occur at
func (sl *SkipList) AllUnique() iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		if sl == nil {
			return
		}
		cur := sl.head
		top := sl.height - 1

		seen := make(map[int]bool)
		var ok bool

		for i := top; i >= 0; i-- {
			for cur.next[i] != nil { // loop through bottom row only
				cur = cur.next[i]

				// track if value has been seen
				if _, ok = seen[cur.value]; !ok {
					seen[cur.value] = true

					// don't yield if cur.value has been previously seen
					if !yield(i, cur.value) {
						return
					}
				}
			}

			cur = sl.head
		}
	}
}

// Return iterator showing search path to find target value
func (sl *SkipList) PathTraverse(target int) iter.Seq2[int, int] {
	if sl == nil {
		return nil
	}
	iter := sl.AllUnique()
	iter = filter(iter, lessThanOrEqual, target)
	return iter
}

func lessThanOrEqual(value int, acc int, target int) (int, bool) {
	// if the value is the highest value that is lower than the target, update acc
	if value <= target && value > acc {
		acc = value
		return acc, true // include new value
	} else {
		return acc, false // don't include new value
	}
}

func filter(it iter.Seq2[int, int], keep func(int, int, int) (int, bool), target int) iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		var ok bool
		acc := 0
		for k, v := range it {
			if acc, ok = keep(v, acc, target); ok {
				if !yield(k, v) {
					break
				}
			}
		}
	}
}
