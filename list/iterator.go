package list

import (
        "iter"
)

// Go 1.23 (released a month ago) now supports standardized iterators, hence add them
// Values is an iterator over the unique elements of sl
func (sl *SkipList) Values() iter.Seq[int] {
        return func(yield func(int) bool) {
                cur := sl.head
                for cur.next[0] != nil { // loop through bottom row only
                        cur = cur.next[0]
                        if !yield(cur.value) {
                                return
                        }
                }
        }
}

// All is an iterator over the elements of sl, including level info
func (sl *SkipList) All() iter.Seq2[int, int] {
        return func(yield func(int, int) bool) {
                cur := sl.head
                top := sl.height -1

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

// All is an iterator over the elements of sl, including level info
func (sl *SkipList) AllUnique() iter.Seq2[int, int] {
        return func(yield func(int, int) bool) {
                cur := sl.head
                top := sl.height -1

                seen := make(map[int]bool)
                var ok bool

                for i := top; i >= 0; i-- {
                        for cur.next[i] != nil { // loop through bottom row only
                                cur = cur.next[i]

                                // track if value has been seen
                                if _, ok = seen[cur.value]; !ok {
                                        seen[cur.value] = true
                                }

                                // don't yield if cur.value has been previously seen
                                if !ok && !yield(i, cur.value) {
                                        return
                                }
                        }

                        cur = sl.head
                }
        }
}


func  (sl *SkipList) PathTraverse(target int) iter.Seq2[int, int] {
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
