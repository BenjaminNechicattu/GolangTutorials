package main
import "fmt"
func main() {
    var a = 10
    var b = 10

    fmt.Println("Go Ladder If...else Statement")

    if (a > b) {
        fmt.Printf("a is greater than b")
    } else if(a == b){
        fmt.Printf("a and b are equal")
    } else{
        fmt.Printf("b is greater than a")
    }
}
