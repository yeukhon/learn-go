package main

import "fmt"

func printMedian(scores []int, score int) {
    fmt.Println("Median of", scores, "is", score)
}

func main() {
  scores := []int{
    10,
    20,
    30,
    40,
//    50,
  }

  var even bool = len(scores) % 2 == 0
  var mid int = len(scores) / 2

  if len(scores) == 0 {
    fmt.Println("No scores available")
  } else if len (scores) == 1 {
    printMedian(scores, scores[0])
  } else if even == false {
    printMedian(scores, scores[mid])
  } else {
    printMedian(scores, (scores[mid-1] + scores[mid])/2)
  }
}
