/* syntax
for <condition1> {
    // loop body
    if (condition2) {
        continue
    }
    // loop body
}
*/

package main
import "fmt"

func main() {

	var ctr = 0
	
	fmt.Println("Go continue statement.")
	
	for ctr < 10 {
		ctr++
		if (ctr == 5) {
            println("5 is skipped")
            continue
            println("This won't be printed too.")
        }
        fmt.Printf("Number is %d\n", ctr)

	}
	fmt.Println("Out of loop")

}
