// NOTE: For now skip generics, everything is an int type

package list

import (
	"fmt"
        "cmp"
)

// SkipList operation type enum
type OpType int

const (
	Insert OpType = iota
	Delete
	Exists
)

const (
	SkipLevels = 5
)

type SkipList[T cmp.Ordered] struct {
	head   *Node[T]
	height int
}

// Constructor
func NewList[T cmp.Ordered]() *SkipList[T] {
	var n Node[T]
	n.next = make([]*Node[T], SkipLevels)
	return &SkipList[T]{head: &n, height: 1} // always have a head node so height is always 1
}

// Simple find which indicates whether value was found and returns relevant node
func (sl *SkipList[T]) Find(value T) (*Node[T], bool) {
	if sl == nil {
		return nil, false
	}
	node, _, _ := sl.FindWithOp(value, Exists)
	return node, node != nil
}

// Find node with specified value given SkipList operation type
func (sl *SkipList[T]) FindWithOp(value T, opType OpType) (*Node[T], []*Node[T], int) {
	if sl == nil {
		return nil, nil, -1
	}
	cur, top, startLevel := sl.head, sl.height-1, -1
	var node *Node[T]
	var prevs []*Node[T]

Outer: // start top down
	for i := top; i >= 0; i-- {
		for cur.next[i] != nil {
			if cur.next[i].value < value {
				cur = cur.next[i]
			} else if cur.next[i].value == value {
				if len(prevs) == 0 { // guard to run only once
					startLevel = i
					node = cur.next[i]
				}

				switch opType {
				case Insert:
					prevs = nil
					break Outer // value already exists!
				case Delete:
					// store previous node to found node at each contiguous level found
					prevs = append(prevs, cur)
				default:
				}
				break
			} else {
				// cur.next value greater than value, hence stay at cur node but drop down a level
				break
			}
		}

		if opType == Insert {
			// store largest node (cur) that is smaller than value at each level
			prevs = append(prevs, cur)
			if startLevel == -1 {
				startLevel = i
			}
		}
	}

	return node, prevs, startLevel
}

// Add new value if not already present to SkipList
func (sl *SkipList[T]) Add(value T) {
	if sl == nil {
		return
	}
	node, nodeHeight := NewNode(value)
	fmt.Printf("Adding %v to %v levels from floor", value, nodeHeight)

	if nodeHeight > sl.height {
		sl.height = nodeHeight
	}

	var prev *Node[T]
	found, prevs, level := sl.FindWithOp(value, Insert)

	if found != nil {
		fmt.Printf("Value %v already found, hence no add", value)
		return
	}

	for i := level; i >= 0; i-- { // Traverse vertically from top to bottom
		if i <= nodeHeight-1 { // Link new node as long as it matches new node's levels
			prev = prevs[level-i]       // Select prev node given each level where node value highest-lower than value is found
			node.next[i] = prev.next[i] // New node linked after prev node
			prev.next[i] = node
		}
	}
}

func (sl *SkipList[T]) Delete(value T) bool {
	if sl == nil {
		return false
	}
	var prev *Node[T]
	del, prevs, level := sl.FindWithOp(value, Delete)

	if del == nil {
		fmt.Printf("Value %v not found, hence no deletion", value)
		return false
	}

	for i := level; i >= 0; i-- { // Traverse vertically from start level (where node found) to bottom
		prev = prevs[level-i] // Select prev node given each level where delete node is found
		prev.next[i] = del.next[i]
	}

	sl.Prune()
	return true
}

// Check for vertical levels that are empty, given that a node has been deleted
// Empty levels can only be singular or contiguous
func (sl *SkipList[T]) Prune() bool {
	if sl == nil {
		return false
	}
	flag := false // indicates whether a vertical prune was done
	h := sl.height - 1

	for i := h; i > 0; i-- { // ignore base level 0
		if sl.head.next[i] == nil {
			flag = true
			h--
		}
	}

	sl.height = h + 1
	return flag
}

// Display skiplist structure ensuring columns are aligned
func (sl *SkipList[T]) Display() {
	if sl == nil {
		return
	}
	fmt.Println("\nSkip List:")
	cur := sl.head

	// store T->column info in map
	columns := make(map[T]int)

	// count elements in botttom row
	for i := 0; cur.next[0] != nil; i++ {
		cur = cur.next[0]
		columns[cur.value] = i
	}

	sl.Show(columns)
}

// Show skiplist structure with or without aligning columns
func (sl *SkipList[T]) Show(columns map[T]int) {
	if sl == nil {
		return
	}
	top := sl.height - 1
	var value T

	for i := top; i >= 0; i-- {
		// show all the node values from the head node,
		// hence reset back to head for each level iteration
		cur := sl.head
                fmt.Printf("L%02d ", i)

		for col := 0; cur.next[i] != nil; col++ {
			value = cur.next[i].value

			// print out arrow base until column matches correct columns value
			for columns != nil {
				if columns[value] == col {
                                        fmt.Printf("-> %02v ", value)
					break
				} else {
					fmt.Printf("------") // extend arrow base
					col++
				}
			}

			if columns == nil {
				fmt.Printf("-> %02d ", value)
			}

			cur = cur.next[i]
		}

		fmt.Printf("-> nil \n")
	}

	fmt.Println("")
}
