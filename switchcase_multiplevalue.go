package main
import "fmt"
func main() {

	var dayOfWeek = 5
	fmt.Println("W3Adda - Go switch with multiple case combnied.")
	switch dayOfWeek {
		case 1, 2, 3, 4, 5:
			fmt.Println("It's a Weekday.")
		case 6, 7:
			fmt.Println("Its Weekend.")
		default:
			fmt.Println("Invalid Day")		
	}

}
