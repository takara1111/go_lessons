package main

import "fmt"

// func hi(name string) string {
// 	// fmt.Println("hi!" + name)
// 	msg := "hi!" + name
// 	return msg
// }

func hi(name string) (msg string) {
	// fmt.Println("hi!" + name)
	msg = "hi!" + name
	return 
}

func main() {
	// hi("takara")
	fmt.Println(hi("takara"))
}