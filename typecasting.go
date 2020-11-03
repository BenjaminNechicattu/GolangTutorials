/*typecasting syntax :

  type_name(expression)
  
*/

package main

import "fmt"

func main() {
   var sum int = 17                           //sum is int
   var count int = 5                          //count is int
   var mean float32                           // mean is float32
   
   mean = float32(sum)/float32(count)         //sum and count are type casted to float32
   fmt.Printf("Value of mean : %f\n",mean)
}
