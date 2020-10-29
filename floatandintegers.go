package main
import "fmt"

func main() {
	var num1 float32 = 10.1514
	var num2 = 5014.645 // Type inferred as float64
	var num = 100
	fmt.Printf("num1 = %f and \n num2 = %f\n also num = %d", num1, num2, num)
}
