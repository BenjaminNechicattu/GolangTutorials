package main
import "fmt"

func main() {
	var marks = 55 
	fmt.Println("Go switch with no expression.")
	switch {
		case marks > 60:                        // boolean expresion
			fmt.Println("Passed with 'A' Grade")
		case marks < 60 && marks >=50:          // boolean expresion
			fmt.Println("Passed with 'B' Grade")
		case marks < 50 && marks >= 35:
			fmt.Println("Passed with 'C' Grade")
		default:
			fmt.Println("You're Fail")	
	}
}
