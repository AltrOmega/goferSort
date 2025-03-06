package main

import (
  "fmt"
  "goferSort/algorithms"
  "goferSort/dataGen"
  "os"
  "github.com/akamensky/argparse"
)




func main() {

  // Todo: give proper usage descriptions
  parser := argparse.NewParser("sort", "insert description here")
  
  bubble := parser.NewCommand("bubble", "bubble sort description")
  merge := parser.NewCommand("merge", "merge sort description")
  quick := parser.NewCommand("quick", "quick sort description")

  print_lvl := parser.Int("p", "print", &argparse.Options{Help: "print level"})
  range_ := parser.Int("r", "range", &argparse.Options{Help: "0+- range overrides minum and maximum"})
  minimum := parser.Int("m", "minimum", &argparse.Options{Help: "lowest numer"})
  maximum := parser.Int("x", "maximum", &argparse.Options{Help: "bigest numer"})
  size := parser.Int("s", "size", &argparse.Options{Help: "number of datapoints"})



  err := parser.Parse(os.Args)
  if err != nil {
    fmt.Println(parser.Usage(err))
    os.Exit(1)
  }

  if *range_ != 0{
    *minimum = int(0-*range_)
    *maximum = int(*range_)
  }
  if *minimum == 0 {
    *minimum = -10
  }
  if *maximum == 0 {
    *maximum = 10
  }
  if *size == 0 {
    *size = 10
  }
  
  // Todo: add error handling to dataGet so it gives one when min > max
  data := dataGen.GenerateRandomSlice(*size, *minimum, *maximum)
  
  // Todo: make it use print levels instead of bool
  do_print := false
  if *print_lvl > 0 {
    do_print = true
  }

  
  fmt.Printf("Data: \n%v\n", data)
  fmt.Printf("Sorting data using: %v\n", os.Args[1])

  if bubble.Happened() {
    data = algorithms.BubbleSort(data, do_print)
  }else if merge.Happened(){
    data = algorithms.MergeSort(data, do_print)
  }else if quick.Happened(){
    data = algorithms.QuickSort(data, 0, len(data)-1, do_print)
  }

  fmt.Printf("Sorted! : \n%v\n", data)
  fmt.Println()
}
