package main

import (
  "fmt"
  "goferSort/algorithms"
  "goferSort/dataGen"
  "os"
  "strconv"
  "errors"
  "strings"
)





func argParse(args []string) (prog_name string, algorithm string, print_lvl int, err error) {
  switch len(args){
  case 2:
    return args[0], args[1], 10, nil

  case 3:
    num, err := strconv.Atoi(args[2])
    if err != nil {
      return "", "", 0, errors.New("Second argument must be a number.")
    }

    return args[0], args[1], num, nil

  default:
    return "", "", 0, errors.New(get_usage_str(prog_name))
  }
}





func get_usage_str(prog_name string) string {
  return fmt.Sprintf("%v <QuickSort | MergeSort | BubbleSort> [number from 0 to 10]", prog_name)
}


func sort_start(algorithm string, data []int) {
  fmt.Printf("Data: \n%v\n", data)
  fmt.Printf("Sorting data using: %v\n", algorithm)
}


func sort_end(data []int) {
  fmt.Printf("Sorted! : \n%v\n", data)
  fmt.Println()
}


func main() {
  prog_name, algorithm, print_lvl, err := argParse(os.Args)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  data := dataGen.GenerateRandomSlice(10, -10, 10)
  
  // Todo: make it use print levels instead of bool
  do_print := false
  if print_lvl > 0 {
    do_print = true
  }

  switch strings.ToLower(algorithm) {
  case "quicksort":
    sort_start(algorithm, data)
    data = algorithms.QuickSort(data, 0, len(data)-1, do_print)
    sort_end(data)
  case "mergesort":
    sort_start(algorithm, data)
    data = algorithms.MergeSort(data, do_print)
    sort_end(data)
  case "bubblesort":
    sort_start(algorithm, data)
    data = algorithms.BubbleSort(data, do_print)
    sort_end(data)
  default:
    fmt.Println("Given algorithm does not exist")
    fmt.Println(get_usage_str(prog_name))
    os.Exit(1)
  }
}
