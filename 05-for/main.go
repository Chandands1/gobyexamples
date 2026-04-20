// Go by Example: ${example#*-} (replace with proper title if you want)
package main

import "fmt"

func main() {
   i :=1
   for i <=3{
    fmt.Println("i=",i)
    i = i +1
   } 

   for j := 0 ; j <=3 ; j++{
    fmt.Println("j=",j)

   }

   for i := range 5{
    fmt.Println("range:", i)
   }

   for n := range 6{
      if n % 2 ==0{
         continue
      }
      fmt.Println(n)
   }
}
