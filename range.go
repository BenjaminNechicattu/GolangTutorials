/*
    for key, value := range collection {
        //block of statements
    }
*/


package main
import "fmt"

func main() {

	langs := []string{"Go", "C", "C++", "Java"}
	fmt.Println("Go For Each Loop.")
	for i,s := range langs{
		fmt.Println(i, s)
	}
}
