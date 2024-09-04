import (
  "fmt"
  "strconv"
)

func main() {
  var sum int
  for i := 0; i < 5; i++ {
	  if i%2 == 0 {
		  fmt.Printf("even: %d", i)
	  }
	  else {
		  fmt.Printf("odd: %d", i)
	  }
  }
}