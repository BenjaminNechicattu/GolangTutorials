// how to get input from the user 
package main 
  
import "fmt"
  
// main function 
func main() { 
  
    // Println function is used to 
    // display output in the next line 
    
    fmt.Println("Enter Your First Name: ") 
    var first string 
    // Taking input from user  using scanln
    fmt.Scanln(&first) 
    
    fmt.Println("Enter Second Last Name: ") 
    var second string 
    fmt.Scanln(&second) 
  
    fmt.Print("Your Full Name is: ") 
  
    // Addition of two string 
    fmt.Print(first + " " + second) 
}
