package main

import (
        "cmp"
        "fmt"
        "github.com/brpandey/skiplist/list"
)

func main() {
        var toggle bool = false

        if !toggle {
                examples := []int{7, 4, 15, 20, 12, 6}
                demo(examples)
        } else {
                examples := []string{"ba", "ze", "ew", "fl", "co", "di"}
                demo(examples)
        }
}

func demo[T cmp.Ordered](examples []T) {
        var notFound T
        sl := list.NewList[T]()

        for i := 0; i < len(examples)-1; i++ {
                v := examples[i]
                fmt.Printf("Add %v\n", v)
                sl.Add(v)
                sl.Display()
        }

        findi2, flag1 := sl.Find(examples[2])
        findNF, flag2 := sl.Find(notFound)
        findi4, flag3 := sl.Find(examples[4])

        fmt.Println("Find ", examples[2], ", found? ", flag1, " value: ", findi2)
        fmt.Println("Find ", notFound, ", found? ", flag2, " value: ", findNF)
        fmt.Println("Find ", examples[4], ", found? ", flag3, " value: ", findi4)

        del := examples[len(examples)-2]

        fmt.Println("")
        fmt.Println("Delete ", del)
        sl.Delete(del)
        sl.Display()

        findDel, flag4 := sl.Find(del)

        fmt.Printf("Find %v (again), found? %v  value: %v\n", del, flag4, findDel)

        // Add last value 
        v := examples[len(examples)-1]

        fmt.Println("")
        fmt.Println("Add ", v)
        sl.Add(v)
        sl.Display()

        // Demo Iterator functionality w/ All and AllWithLevels
        fmt.Printf("Values: ")
        for v := range sl.Values() {
                fmt.Printf("%v ", v)
        }

        fmt.Println("")

        fmt.Printf("All w/ Levels: ")
        var prev int

        for l, v := range sl.All() {
                if l != prev {
                        fmt.Println("")
                        prev = l
                }

                fmt.Printf("(%v, %v) ", l, v)
        }

        fmt.Println("\n")
        fmt.Printf("All Unique by Levels: ")

        for l, v := range sl.AllUnique() {
                fmt.Printf("(%v, %v) ", l, v)
        }

        fmt.Println("\n")

        target := examples[3]
        fmt.Printf("Path Traversal for target %v: ", target)

        var len = 0
        for l, v := range sl.PathTraverse(target) {
                fmt.Printf("(%v, %v) ", l, v)
                len++
        }

        fmt.Printf("\nPath Traversal length is %v\n", len)
}
