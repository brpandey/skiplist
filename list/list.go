// NOTE: For now skip generics, everything is an int type

package list

import (
        "fmt"
        "math/rand"
)

const (
        SkipLevels = 5
)

type SkipList struct {
        head *Node
        height int
}

type Node struct {
        value int
        next [] *Node
}

func NewList() *SkipList {
        var n Node
        n.next = make([]*Node, SkipLevels)

        return &SkipList {head: &n, height: 1}
}

func (sl *SkipList) Add (value int) {
        nodeLevelIndex := RandomLevel()

        fmt.Printf("Adding %d to %d levels from floor", value, nodeLevelIndex+1)

        next := make([]*Node, nodeLevelIndex+1)
        node := &Node{value: value, next: next}

        if nodeLevelIndex+1 > sl.height {
                sl.height = nodeLevelIndex+1
        }

        cur := sl.head
        top := sl.height - 1

        // traverse vertically from top to bottom
Outer:
        for i := top; i >= 0; i-- {
                for cur.next[i] != nil {
                        if cur.next[i].value < value {
                                cur = cur.next[i] // advance along level's linked list
                        } else if cur.next[i].value == value {
                                break Outer // value already exists -> no need to insert
                        } else {
                                // if cur value is greater than value go to next level below
                                // by breaking from inner and increment level with the same cur node
                                break
                        }
                }

                if i <= nodeLevelIndex { // link new node as long as it matches new node's levels
                        // link new node after current node
                        node.next[i] = cur.next[i]
                        cur.next[i] = node
                }
        }
}

func (sl *SkipList) Delete (value int) bool {
        var prev *Node
        del, prevs, level := sl.Find(value)

        if del == nil {
                fmt.Printf("Value %d not found, hence no deletion", value)
                return false
        }

        for i := level; i >= 0; i-- { // Traverse vertically from top to bottom
                prev = prevs[level-i] // select prev node given each level where delete node is found
                prev.next[i] = del.next[i];
        }

        sl.Prune()
        return true
}

// Check for vertical levels that are empty given deleted node
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

func (sl *SkipList) Find (value int) (*Node, []*Node, int) {
        cur, top := sl.head, sl.height - 1
        startLevel := 0

        var node *Node
        var prevs []*Node

        // start top down
        for i := top; i >= 0; i-- {
                for cur.next[i] != nil {
                        if cur.next[i].value < value {
                                cur = cur.next[i]
                        } else if cur.next[i].value == value {
                                if len(prevs) == 0 {
                                        startLevel = i
                                        node = cur.next[i]
                                }
                                // store previous node to found node at each contiguous level found
                                prevs = append(prevs, cur) 
                                break
                        } else {
                                // stay at cur node but drop down a level
                                break
                        }
                }
        }

        if len(prevs) > 0 {
                return node, prevs, startLevel
        } else {
                return nil, nil, -1
        }
}

func (sl *SkipList) Display () {
        fmt.Println("\nSkip List:")
        top := sl.height - 1

        for i := top; i >= 0; i-- {
                // show all the node values from the head node,
                // hence reset back to head for each level
                cur := sl.head
                fmt.Printf("L%d ", i)

                for cur.next[i] != nil {
                        fmt.Printf("%d ---> ", cur.next[i].value)
                        cur = cur.next[i]
                }

                fmt.Printf(" nil \n")
        }

        fmt.Println("")
}

func RandomLevel () int {
        var level int

        // every time the coin flip returns true, track increments
        // as long as they are valid
        for rand.Intn(2) == 0 && level < SkipLevels - 1 {
                level++
        }

        return level
}
