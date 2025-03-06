package algorithms

import "fmt"

func swap(data []int, a, b int, print_lvl bool) []int{
    if print_lvl {
      fmt.Printf("# swap data: %v; a: %v; b: %v;\n", data, a, b)
    }
  temp := data[a]
  data[a] = data[b]
  data[b] = temp
  return data
}

/*
Returns sorted data using quickSort
Todo: add an explanation of how it does what it does
*/
func QuickSort(data []int, start, end int, print_lvl bool) []int {
    if print_lvl {
      fmt.Printf("### QuickSort data: %v; start: %v; end: %v;\n", data, start, end)
    }

    if start >= end {
      if print_lvl {
        fmt.Println("start >= end returning data as is\n")
      }
      return data
    } 

  p := partition(data, start, end, print_lvl)
  QuickSort(data, start, p-1, print_lvl)
  QuickSort(data, p+1, end, print_lvl)
  
  return data
} 

func partition(data []int, start int, end int, print_lvl bool) int{
  pivot := data[end]
  if print_lvl { 
    fmt.Printf("# partition start: %v; end: %v; pivot_value: %v;\n", start, end, pivot)
  }

  i := start-1 
  for j := start; j < end; j++ {
    if data[j] < pivot {
      i++
      swap(data, i, j, print_lvl)
    }
  }

  if print_lvl { 
    fmt.Printf("# final swap\n")
  }
  swap(data, i+1, end, print_lvl)
  return i+1
}





func MergeSort(data []int, print_lvl bool) []int {
  if len(data) < 2{
    return data
  }
  
  pivot := len(data) / 2

  left := MergeSort( data[:pivot], print_lvl)
  right := MergeSort( data[pivot:], print_lvl)

  return merge(left, right, print_lvl)
}

func merge(left []int, right []int, print_lvl bool) []int {
  var sorted []int
  if print_lvl {
    fmt.Println(fmt.Sprintf("# merge on: %v : %v", left, right))
  }

  i := 0
  j := 0
  for i < len(left) && j < len(right) {
    if left[i] <= right[j] {
      sorted = append(sorted, left[i])
      i++
      continue
    }
    sorted = append(sorted, right[j])
    j++
  }
  
  if i < len(left){
  sorted = append(sorted, left[i:]...)
  }
  if j < len(right){
  sorted = append(sorted, right[j:]...)
  }

  if print_lvl {
    fmt.Println(fmt.Sprintf("# result: %v", sorted))
  }
  return sorted
  
}





func BubbleSort(data []int, print_lvl bool) []int {
  for i := 0; i < len(data); i++ {
    for j := 0; j < len(data)-1-i; j++ {
      if data[j] > data[j+1]{
        swap(data, j, j+1, print_lvl)
      }
    }
  }
  return data
}
