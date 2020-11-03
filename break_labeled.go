package main
import "fmt"

func main() {

	var count = 0
	
	fmt.Println("Go labeled break statement.")
	
	outer:
	for count <= 10 {
		count++
		if (count == 5) {
            break outer
        }
        fmt.Printf("Inside loop %d\n", count)
	}
	fmt.Println("Out of loop")

}
