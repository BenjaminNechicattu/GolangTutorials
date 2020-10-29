package main
import "fmt"

func main() {
	var varByte byte = 'a'
	var varRune rune = '♥'
	var  a = 100
	fmt.Printf("%c = %d and %c = %U\n also a = %d", varByte, varByte, varRune, varRune, a)
}
