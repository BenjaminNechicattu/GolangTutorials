package main
import "fmt"
func main() {
  //bitwise operators applicable only on integers to perform bitwise operation to its operands.
  
    var a,b,c int
    a = 50
    b = 10
    
    c = a & b       // &    bitwise AND            integers
    fmt.Println(c)
    
    c = a | b       // |    bitwise OR             integers
    fmt.Println(c)
    
    c = a ^ b       // ^    bitwise aOR            integers
    fmt.Println(c)
    
    c = a &^ b      // &^   bit clear (AND NOT)    integers
    fmt.Println(c)
}
