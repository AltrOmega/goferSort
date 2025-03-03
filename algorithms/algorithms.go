package algorithms

import "fmt"

func swap(data []int, a, b int, do_print bool) []int{
    if do_print {
      fmt.Printf("# swap data: %v; a: %v; b: %v;\n", data, a, b)
    }
  temp := data[a]
  data[a] = data[b]
  data[b] = temp
  return data
}

/*
Return a sorted data using quickSort
Todo: add an explanation of how it does what it does
*/
func QuickSort(data []int, start, end int, do_print bool) []int {
    if do_print {
      fmt.Printf("### QuickSort data: %v; start: %v; end: %v;\n", data, start, end)
    }

    if start >= end {
      if do_print {
        fmt.Println("start >= end returning data as is\n")
      }
      return data
    } 

  p := partition(data, start, end, do_print)
  QuickSort(data, start, p-1, do_print)
  QuickSort(data, p+1, end, do_print)
  
  return data
} 

func partition(data []int, start int, end int, do_print bool) int{
  pivot := data[end]
  if do_print { 
    fmt.Printf("# partition start: %v; end: %v; pivot_value: %v;\n", start, end, pivot)
  }

  i := start-1 
  for j := start; j < end; j++ {
    if data[j] < pivot {
      i++
      swap(data, i, j, do_print)
    }
  }

  if do_print { 
    fmt.Printf("# final swap\n")
  }
  swap(data, i+1, end, do_print)
  return i+1
}
