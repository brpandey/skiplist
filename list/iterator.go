package list

import (
	"iter"
        "cmp"
)

// Go 1.23 supports standardized iterators
// Values is an iterator over the value elements of sl
func (sl *SkipList[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
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
func (sl *SkipList[T]) All() iter.Seq2[int, T] {
        return sl.allWith(nil)
}

// AllUnique returns an iterator over the unique elements of sl with the highest
// level they first occur at
func (sl *SkipList[T]) AllUnique() iter.Seq2[int, T] {
        seen := make(map[T]bool)
        return sl.allWith(seen)
}

// allWith returns an iterator over the elements of sl with the highest
// level they first occur at, if seen is passed in only the unique elements are returned
func (sl *SkipList[T]) allWith(seen map[T]bool) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		if sl == nil {
			return
		}
		cur := sl.head
		top := sl.height - 1

		var ok bool

		for i := top; i >= 0; i-- {
			for cur.next[i] != nil { // loop through bottom row only
				cur = cur.next[i]

                                if seen != nil {
                                        // track if value has been seen
                                        _, ok = seen[cur.value]
                                        seen[cur.value] = true
                                }

                                // 1) don't yield if cur.value has been previously seen
                                // or 2) if seen not specified, attempt to yield
                                if !ok && !yield(i, cur.value) {
                                        return
                                }
			}

			cur = sl.head
		}
	}
}



// Return iterator showing search path to find target value
func (sl *SkipList[T]) PathTraverse(target T) iter.Seq2[int, T] {
	if sl == nil {
		return nil
	}
	iter := sl.AllUnique()
	iter = filter(iter, highestLower, target)
	return iter
}


func highestLower[T cmp.Ordered](value T, acc T, target T) (T, bool) {
	// if the value is the highest value that is lower than the target, update acc
	if value <= target && value > acc {
		acc = value
		return acc, true // include new value
	} else {
		return acc, false // don't include new value
	}
}

func filter[T cmp.Ordered](it iter.Seq2[int, T], keep func(T, T, T) (T, bool), target T) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		var ok bool
		var acc T // needs to be a generic function e.g. min(T) to produce min value for T, hence use positive numbers
		for k, v := range it {
			if acc, ok = keep(v, acc, target); ok {
				if !yield(k, v) {
					break
				}
			}
		}
	}
}
