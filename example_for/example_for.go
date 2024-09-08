package example_for

import (
  "fmt"
  "strconv"
)

func example_for() {
  var sum int
  for i := 0; i < 5; i++ {
    sum += i
      fmt.Println(i)
  }

  fmt.Println("sum: "+ strconv.Itoa(sum))
}