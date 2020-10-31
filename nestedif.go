package main
import "fmt"
func main() {
    var a = 25
    fmt.Println("Go Nested If Statement")

    if (a < 100) {
        if (a < 50) {
           fmt.Println("a is less than 50")
        }

        if (a >= 50) {
           fmt.Println("a is greater than 50")
        }
        
    } 

}
