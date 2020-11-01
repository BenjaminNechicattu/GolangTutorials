/* Syntax :
  for initialization; condition; increment {
    // loop body
  }
*/

package main
import "fmt"

func main() {
	sum := 0
	fmt.Println("Go For loop.")
	for i := 1; i < 5; i++ {
		sum += i
	}
	fmt.Println(sum)
}
