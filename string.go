package main  
import "fmt"  
func main() {  
   // go does not allows newlines, and can contain escape sequences like `\n`, `\t` etc.
	var fstr="Hello World"

	// Can span multiple lines. Escape characters are not allowed.
	var sstr= `Hello world, this
	a multi-line text string.`
	fmt.Println(fstr)  
	fmt.Println(sstr)  
  
  // run and see the difference
}
