package main
import "fmt"

func main() {

	fmt.Println("Go Range with Strings")
	for i,s := range "My Name is Benjamin G Nechicattu"{
		fmt.Printf("%U represents %c and it is at position %d\n", s, s, i)
	}
}
