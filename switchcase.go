/* Syntax :

switch <expression>{
case <val1>:
  //Block of code to be executed if expression = val1;
case <val2>:
  //Block of code to be executed if expression = val2;
default:
  //Optional
  //Block of code to be executed
  //if expression is different 
  //from both val1 and val2;
}

*/

package main
import "fmt"
func main() {

	var dayOfWeek = 5
	fmt.Println("Go switch statement.")
  
	switch dayOfWeek {
	case 1:
		fmt.Println("Today is Monday.")
	case 2:
		fmt.Println("Today is Tuesday.")
	case 3:
		fmt.Println("Today is Wednesday.")
	case 4:
		fmt.Println("Today is Thursday.")		
	case 5:
		fmt.Println("Today is Friday.")
	case 6:
		fmt.Println("Today is Saturday.")	
	case 7:
		fmt.Println("Today is Sunday.")	
	default:
		fmt.Println("Invalid Weekday.")
	}

}
