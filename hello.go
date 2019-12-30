package main

import "fmt"

func main()  {
	// const name = "takara"
	// name = "tanaka"
	// fmt.Println

	const (
		sun = iota
		mon 
		tue
	)
	fmt.Println(sun, mon ,tue)
}