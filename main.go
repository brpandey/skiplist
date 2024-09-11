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

        flag1, find15 := sl.Find(15)
        flag2, find10 := sl.Find(10)

        fmt.Println("Find 15, found? ", flag1, " value: ", find15)
        fmt.Println("Find 10, found? ", flag2, " value: ", find10)
  }
