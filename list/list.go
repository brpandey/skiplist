// NOTE: For now skip generics, everything is an int type

package list

import (
        "fmt"
        "math/rand"
        "iter"
)

type Node struct {
        value int
        next [] *Node
}

func NewNode(value int) (*Node, int) {
        height := RandomLevel() + 1
        next := make([]*Node, height)
        node := &Node{value: value, next: next}
        return node, height
}

// OpType enum
type OpType int

const (
	Insert OpType = iota
        Delete
        Exists
)

const (
        SkipLevels = 5
)

type SkipList struct {
        head *Node
        height int
}

// Constructor
func NewList() *SkipList {
        var n Node
        n.next = make([]*Node, SkipLevels)
        return &SkipList {head: &n, height: 1}
}

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


// Simple find which indicates whether value was found and returns relevant node
func (sl *SkipList) Find (value int) (*Node, bool) {
        node, _, _ := sl.FindWithOp(value, Exists)
        return node, node != nil
}

// Find node with specified value given SkipList operation type
func (sl *SkipList) FindWithOp (value int, opType OpType) (*Node, []*Node, int) {
        cur, top, startLevel := sl.head, sl.height - 1, -1
        var node *Node
        var prevs []*Node

Outer:         // start top down
        for i := top; i >= 0; i-- {
                for cur.next[i] != nil {
                        if cur.next[i].value < value {
                                cur = cur.next[i]
                        } else if cur.next[i].value == value {
                                if len(prevs) == 0 { // guard to run only once
                                        startLevel = i
                                        node = cur.next[i]
                                }

                                if opType == Insert {
                                        prevs = nil
                                        break Outer // value already exists!
                                }

                                if opType == Delete {
                                        // store previous node to found node at each contiguous level found
                                        prevs = append(prevs, cur)
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
                        if startLevel == -1 { startLevel = i }
                }
        }

        return node, prevs, startLevel
}

// Add new value if not already present to SkipList
func (sl *SkipList) Add (value int) {
        node, nodeHeight := NewNode(value)
        fmt.Printf("Adding %d to %d levels from floor", value, nodeHeight)

        if nodeHeight > sl.height {
                sl.height = nodeHeight
        }

        var prev *Node
        found, prevs, level := sl.FindWithOp(value, Insert)

        if found != nil {
                fmt.Printf("Value %d already found, hence no add", value)
                return
        }

        for i := level; i >= 0; i-- { // Traverse vertically from top to bottom
                if i <= nodeHeight-1 { // Link new node as long as it matches new node's levels
                        prev = prevs[level-i] // Select prev node given each level where node value highest-lower than value is found
                        node.next[i] = prev.next[i] // New node linked after prev node
                        prev.next[i] = node
                }
        }
}

func (sl *SkipList) Delete (value int) bool {
        var prev *Node
        del, prevs, level := sl.FindWithOp(value, Delete)

        if del == nil {
                fmt.Printf("Value %d not found, hence no deletion", value)
                return false
        }

        for i := level; i >= 0; i-- { // Traverse vertically from start level (where node found) to bottom
                prev = prevs[level-i] // Select prev node given each level where delete node is found
                prev.next[i] = del.next[i];
        }

        sl.Prune()
        return true
}

// Check for vertical levels that are empty given a node has been deleted
// Empty levels can only be singular or contiguous
func (sl *SkipList) Prune () bool {
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
func (sl *SkipList) Display () {
        fmt.Println("\nSkip List:")
        cur := sl.head

        // store key->column info in map
        columns := make(map[int]int)

        // count elements in botttom row
        for i := 0; cur.next[0] != nil; i++ {
                cur = cur.next[0]
                columns[cur.value] = i
        }

        sl.Show(columns)
}

// Show skiplist structure with or without reflecting columns info
func (sl *SkipList) Show (columns map[int]int) {
        top := sl.height - 1
        var value int

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
                                        fmt.Printf("-> %02d ", value)
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

func RandomLevel () int {
        var level int

        // Every time the proverbial "coin flip" returns true, track increments
        // as long as they are valid (1/2 * 1/2 * ....)
        for rand.Intn(2) == 0 && level < SkipLevels - 1 {
                level++
        }

        return level
}
