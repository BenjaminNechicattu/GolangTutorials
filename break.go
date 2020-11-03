package main
import "fmt"

func main() {

	var count = 0
	fmt.Println("Go break statement.")
	for count <= 10 {
		count++
		if (count == 5) {
            break       //breaks innermost loop/switch/select
        }
        fmt.Printf("Inside loop %d\n", count)
	}
	fmt.Println("Out of loop")

}
