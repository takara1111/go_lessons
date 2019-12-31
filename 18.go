package main 

import "fmt"

type user struct {
	name string
	score int
}

func main() {
	// u := new(user)
	// u.name = "takara"
	// u.score = 100
	u := user{"takara", 50}
	fmt.Println(u)
}