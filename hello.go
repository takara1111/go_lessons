package main

import "fmt"

// func swap(a, b int) (int, int) {
// 	return b, a
// }

func main() {
	// f := func(a, b int) (int, int) {
	// 	return b, a
	// }
	
	func(msg string) {
		fmt.Println(msg)
	}("takara")
}