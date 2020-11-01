/*
Refer Maps to know more on maps
*/

package main
import "fmt"

func main() {

	fmt.Println("Go For Each with Map")
  
	fruits := map[string]string{"A": "Apple", "B": "banana", "C": "Cherry"}
  
    for key, value := range fruits {
        fmt.Printf("%s -> %s\n", key, value)
    }

    for key := range fruits {
        fmt.Println("key:", key)
    }
}
