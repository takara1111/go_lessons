package main 

import "fmt"

func main() {
	// m := make(map[string]int)
	// m["takara"] = 200
	// m["satou"] = 300
	m := map[string]int{"takara":100, "satou": 200}
	fmt.Println(m)
	fmt.Println(len(m))
	delete(m, "takara")
	fmt.Println(m)
	v, ok := m["satou"]
	fmt.Println(v)
	fmt.Println(ok)

}