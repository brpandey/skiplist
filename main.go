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

        find15, _, levelA := sl.Find(15)
        find10, _, _ := sl.Find(10)
        find12, _, _ := sl.Find(12)

        fmt.Println("Find 15, found? ", find15 != nil, " value: ", find15, " level: ", levelA)
        fmt.Println("Find 10, found? ", find10 != nil, " value: ", find10)
        fmt.Println("Find 12, found? ", find12 != nil, " value: ", find12)

        fmt.Println("")
        fmt.Println("Delete 12")
        sl.Delete(12)
        sl.Display()

        find12, _, _ = sl.Find(12)

        fmt.Println("Find 12 (again), found? ", find12 != nil, " value: ", find12)
  }
