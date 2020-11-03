package main
import "fmt"

func main() {

	var i = 0
	var j = 0
	
	fmt.Println("Go continue statement.")
	
	    outerfor:
	    for i < 3 {
	    	j = 0
	    	i++
	        for j < 3 {
	        	j++
	            fmt.Printf("i is %d and j is %d\n",i,j)
	            if(i == 2){
	                fmt.Println("I am Skipped")
	                continue outerfor
	            }
	            fmt.Println("I am Printed")
	        }
	        
	    }
	

}
