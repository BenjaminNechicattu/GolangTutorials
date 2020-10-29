package main
import "fmt"

func main() {
	var varByte byte = 'a'
	var varRune rune = '🎞'

	fmt.Printf("%c = %d and %c = %U\n", varByte, varByte, varRune, varRune)
}
