/*
  removing initializationa and incrementation, for loop acts like a while loop
*/


package main
import "fmt"

func main() {
	power := 1
	fmt.Println("Go For loop as while loop.")
	for power < 5 {
		power *= 2
	}
	fmt.Println(power)
}
