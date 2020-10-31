package main
import "fmt"
func main() {
    var a = 10
    var b = 20
    
    /* Syntax :
    if (condition) {
          // statement(s) to be executed when condition is true
                   }
    else {
          // statement(s) to be executed when condition is false
         }
    */
    
    fmt.Println("W3Adda - Go If...else Statement")

    if (a > b) {
        fmt.Printf("a is greater than b")
    } else {
        fmt.Printf("b is greater than a")
    }
}
