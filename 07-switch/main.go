// Go by Example: ${example#*-} (replace with proper title if you want)
package main

import "fmt"
import "time"

func main() {
  i := 2
  
  switch i{
  case 1:
    fmt.Println("ond")
  case 2:
    fmt.Println("two")
  case 3:
    fmt.Println("Three")
  default:
    fmt.Println("Nothing matches")
  }

 
  t := time.Now()

  switch {
  case  t.Hour() < 12:
    fmt.Println("it is Moring")
  default:
    fmt.Println("It is after noon")

  }


  whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}
