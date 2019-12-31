package main 

import "fmt"

func main() {
	// var a [5] int // a[0] - a[4]
	// a[2] = 2
	// a[4] = 10	
	// fmt.Println(b)

	b := [...]int{1, 3, 5, 7} 
	fmt.Println(len(b))

}