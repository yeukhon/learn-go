package main

import "fmt"

// function signature
// func FunctionName(var_name var_type, ...) return_type {}
func isEven(num int) bool {
  if num % 2 == 0 {
    return true
  }
  return false
}

func main() {
  for i:= 1; i <= 10; i++ {
    if isEven(i) {
      fmt.Println(i, "is even")
    } else {
      fmt.Println(i, "is odd")
    }
  }
}
