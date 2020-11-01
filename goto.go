package main  
import "fmt"  
func main() {  
   var age int  
   
election:  //<.........................................................
   fmt.Print("Enter your age ")                                     //.
   fmt.Scanln(&age)                                                 //.
   if (age <= 17) {                                                 //.
      fmt.Print("You are not eligible to vote.\n")                  //.
      goto election  //>...............................................
   } else {  
      fmt.Print("You are eligible to vote.")  
   }  
}
