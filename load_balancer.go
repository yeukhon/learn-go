package main

import "fmt"

// A generator for load balancer.

func weightlessLB(lb_count int) func () int {
  lb_index := 0
  return func() int{
    // A weightless load balancer will return the next lb server
    // which essentially is a (clock) modular arithemtic problem.
    next_lb := lb_index % lb_count
    lb_index += 1
    return next_lb
  }
}

func main() {
  nextLB := weightlessLB(12)
  for i := 1; i <= 24; i++ {
   fmt.Println("Client", i, "connected to", nextLB())
 }
}
