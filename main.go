package main

import (
        "fmt"
        "github.com/brpandey/skiplist/list"
)


func main() {
        sl := list.NewList()

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
                fmt.Printf("%d ", v)
        }

        fmt.Println("")

        fmt.Printf("All w/ Levels: ")
        var prev int

        for l, v := range sl.All() {
                if l != prev {
                        fmt.Println("")
                        prev = l
                }

                fmt.Printf("(%d, %d) ", l, v)
        }

        fmt.Println("")
  }
