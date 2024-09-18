package main

import (
        "fmt"
        "github.com/brpandey/skiplist/list"
)


func main() {
        //skipListInts()
        skipListStrings()
}

func skipListInts() {
        sl := list.NewList[int]()

        fmt.Println("Add 4")
        sl.Add(4)
        sl.Display()

        fmt.Println("Add 7")
        sl.Add(7)
        sl.Display()

        fmt.Println("Add 15")
        sl.Add(15)
        sl.Display()

        fmt.Println("Add 20")
        sl.Add(20)
        sl.Display()

        fmt.Println("Add 12")
        sl.Add(12)
        sl.Display()

        find15, flag1 := sl.Find(15)
        find10, flag2 := sl.Find(10)
        find12, flag3 := sl.Find(12)

        fmt.Println("Find 15, found? ", flag1, " value: ", find15)
        fmt.Println("Find 10, found? ", flag2, " value: ", find10)
        fmt.Println("Find 12, found? ", flag3, " value: ", find12)

        fmt.Println("")
        fmt.Println("Delete 12")
        sl.Delete(12)
        sl.Display()

        find12, flag3 = sl.Find(12)

        fmt.Println("Find 12 (again), found? ", flag3, " value: ", find12)

        fmt.Println("")
        fmt.Println("Add 6")
        sl.Add(6)
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

        target := 15
        fmt.Printf("Path Traversal for target %v: ", target)

        var len = 0
        for l, v := range sl.PathTraverse(target) {
                fmt.Printf("(%v, %v) ", l, v)
                len++
        }

        fmt.Printf("\nPath Traversal length is %v\n", len)
}


func skipListStrings() {
        sl := list.NewList[string]()

        fmt.Println("Add ba")
        sl.Add("ba")
        sl.Display()

        fmt.Println("Add ze")
        sl.Add("ze")
        sl.Display()

        fmt.Println("Add ca")
        sl.Add("ca")
        sl.Display()

        fmt.Println("Add co")
        sl.Add("co")
        sl.Display()

        fmt.Println("Add ew")
        sl.Add("ew")
        sl.Display()

        find1, flag1 := sl.Find("co")
        find2, flag2 := sl.Find("ca")
        find3, flag3 := sl.Find("ca")

        fmt.Println("Find 1, found? ", flag1, " value: ", find1)
        fmt.Println("Find 2, found? ", flag2, " value: ", find2)
        fmt.Println("Find 3, found? ", flag3, " value: ", find3)

        fmt.Println("")
        fmt.Println("Delete co")
        sl.Delete("co")
        sl.Display()

        find1, flag3 = sl.Find("co")

        fmt.Println("Find co (again), found? ", flag3, " value: ", find1)

        fmt.Println("")
        fmt.Println("Add bl")
        sl.Add("bl")
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

        target := "ze"
        fmt.Printf("Path Traversal for target %v: ", target)

        var len = 0
        for l, v := range sl.PathTraverse(target) {
                fmt.Printf("(%v, %v) ", l, v)
                len++
        }

        fmt.Printf("\nPath Traversal length is %v\n", len)
  }
