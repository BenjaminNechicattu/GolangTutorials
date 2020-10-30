package main
import "fmt"
func main() {
    var a int =20
    var b int =10
    var c int =25
    
    var flag bool = false
    var result bool
    //syntax var <name> <type> = <expression>
    
    fmt.Println("W3Adda - Go Logical Operators")
    result = (a > b) && (a > c) // && and
    
    fmt.Printf("(a>b) && (a>c) : %t\n",result)
    result = (a > b) || (a > c) // || or
    
    fmt.Printf("(a>b) || (a>c) :%t\n",result)
    result = !flag // ! not
    
    fmt.Printf("!flag :%t\n",result)
     
}
