package main

import (
  "fmt"
  "goferSort/algorithms"
  "goferSort/dataGen"
)

func main() {
  data := dataGen.GenerateRandomSlice(100, -1000, 1000)
  fmt.Printf("Sorting. : %v\n", data)
  algorithms.QuickSort(data, 0, len(data)-1, false)
  fmt.Printf("Sorted! : %v\n", data)
  fmt.Println()
}
