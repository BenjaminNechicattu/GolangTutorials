package main
import "fmt"
func main() {
    var x = 25
    
    /* syntax : 
    if(condation){  
    statement(s) to be executed when condition is true
    }
    */
    
    fmt.Println("Go If Statement")
    if(x > 10) {
        fmt.Printf("%d is greater than 10\n", x)
    }
}
