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

        fmt.Println("0) Level Index is ", nodeLevelIndex)

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
                                // if cur value is greater go to next level
                                // given the outer for loop iteration with the same cur
                                break
                        }
                }

                if i <= nodeLevelIndex { // link new node as long as it matches new node's levels
                        // link new node after current node
                        node.next[i] = cur.next[i]
                        cur.next[i] = node

                        fmt.Println("3) Linked new Node. head is ", sl.head, "new node is ", node)
                }
        }
}

func (sl *SkipList) Find (value int) (bool, *Node) {
        cur := sl.head

        for i := sl.height - 1; i >= 0; i-- {
                for cur.next[i] != nil {
                        if cur.next[i].value > value {
                                // stay at cur node but just drop down a level on outer loop iteration
                                break
                        }

                        if cur.next[i].value == value {
                                return true, cur.next[i]
                        }

                        cur = cur.next[i]
                        fmt.Println("cur is ", cur)
                }
        }

        return false, nil
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
