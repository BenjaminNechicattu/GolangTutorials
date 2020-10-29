package main
import "fmt"

func main() {
	/* Go doesn't have char data type istead uses Byte and Rune */
	var varByte byte = 'a' /* Byte represents ASCII value */
	var varRune rune = '❤' /* Rune represents Unicode in UTF-8 Format value */
	fmt.Printf("%c = %d and %c = %U\n also a = %d", varByte, varByte, varRune, varRune, a)
	/* %c of varByte is a */
	/* %d of varByte is 97 */
	/* %c of varRune is ❤ */
	/* %U of varRune is U+1F39E */
	/* Run and see differences */
	
}
