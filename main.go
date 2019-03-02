package main

import (
  "flag"
  "fmt"
  "os"
  "strconv"
)

func fibonacciA(n int) int {
  if n < 2 {
    return n
  }

  return fibonacciA(n-2) + fibonacciA(n-1)
}

func fibonacciB(n int) int {
  if n < 2 {
    return n
  }

  ch0 := make(chan int)
  ch1 := make(chan int)

  go func () {
    ch0 <- fibonacciB(n-2)
  }()
  go func () {
    ch1 <- fibonacciB(n-1)
  }()
  f0, f1 := <- ch0, <-ch1

  return f0 + f1
}

func main() {
  flag.Parse()

  n, err := strconv.Atoi(flag.Arg(0))
  if err != nil {
    fmt.Fprintln(os.Stderr, "Cancel: Set number argument.")
    return
  }

  fmt.Println(fibonacciA(n))
  fmt.Println(fibonacciB(n))
}
